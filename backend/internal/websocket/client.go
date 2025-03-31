package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
)

type ChatSession struct {
	clients   map[*websocket.Conn]ClientInfo
	colors    map[string]string
	broadcast chan Message
	mutex     sync.Mutex
	history   []Message
}

type ClientInfo struct {
	UserID     string
	ClientType string
	Nickname   string
}

type ChatMessage struct {
	Type      string          `json:"type"`
	Text      string          `json:"text"`
	From      json.RawMessage `json:"from"`
	To        json.RawMessage `json:"to"`
	UserID    string          `json:"userId"`
	Nickname  string          `json:"nickname"`
	SessionID string          `json:"sessionId"`
	History   []ChatMessage   `json:"history,omitempty"`
}

var chatSessions = make(map[uuid.UUID]*ChatSession)

// Client represents the websocket client at the server
type Client struct {
	// The actual websocket connection.
	conn *websocket.Conn
	send chan []byte
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}
}

func (client *Client) readPump() {
	defer func() {
		client.disconnect()
		log.Println("readPump finished for client")
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error in readPump: %v", err)
			} else {
				log.Printf("Error reading message in readPump: %v", err)
			}
			break
		}

		var message Message
		if err := json.Unmarshal(jsonMessage, &message); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}
		broadcast <- jsonMessage
	}
}

func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for range n {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				log.Printf("Error closing writer in writePump: %v", err)
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Error sending ping in writePump: %v", err)
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	unregister <- client
	close(client.send)
	client.conn.Close()
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	sessionIDStr := r.URL.Query().Get("session_id")
	if sessionIDStr == "" {
		log.Println("session_id is required")
		http.Error(w, "session_id is required", http.StatusBadRequest)
		return
	}

	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		log.Printf("Invalid session_id value: %v", err)
		http.Error(w, "Invalid session_id value", http.StatusBadRequest)
		return
	}

	chatSession, ok := chatSessions[sessionID]
	if !ok {
		log.Printf("Creating new chat session for session_id: %s", sessionID)
		chatSession = &ChatSession{
			clients:   make(map[*websocket.Conn]ClientInfo),
			colors:    make(map[string]string),
			broadcast: make(chan Message),
			history:   []Message{},
		}
		chatSessions[sessionID] = chatSession

		go chatSession.HandleMessages()
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	log.Println("WebSocket connection upgraded successfully")

	var initialMessage ChatMessage
	if err := conn.ReadJSON(&initialMessage); err != nil {
		log.Printf("Error reading initial message: %v", err)
		conn.Close()
		return
	}

	log.Printf("Received initial message: %+v", initialMessage)

	clientType := "editor"
	if r.URL.Path == "/chat" {
		clientType = "chat"
	}

	chatSession.mutex.Lock()
	chatSession.clients[conn] = ClientInfo{
		UserID:     initialMessage.UserID,
		ClientType: clientType,
		Nickname:   initialMessage.Nickname,
	}
	chatSession.colors[initialMessage.UserID] = generateColor()

	historyMessage := Message{
		Type:    "history",
		History: chatSession.history,
	}
	log.Printf("Sending history message: %+v", historyMessage)
	if err := conn.WriteJSON(historyMessage); err != nil {
		log.Printf("Error sending history message: %v", err)
		chatSession.mutex.Unlock()
		conn.Close()
		return
	}

	chatSession.mutex.Unlock()

	client := newClient(conn)

	go client.writePump()
	go client.readPump()

	register <- client
}

func (chatSsn *ChatSession) HandleMessages() {
	for message := range chatSsn.broadcast {
		chatSsn.mutex.Lock()
		var senderNickname string
		for _, info := range chatSsn.clients {
			if info.UserID == message.UserID {
				senderNickname = info.Nickname
				break
			}
		}
		message.Nickname = senderNickname
		if message.Type == "chat" {
			chatSsn.history = append(chatSsn.history, message)
		}
		for client, info := range chatSsn.clients {
			if info.UserID != message.UserID {
				if message.Type == "update" && info.ClientType != "editor" {
					continue
				}
				if message.Type == "chat" && info.ClientType != "chat" {
					continue
				}
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("Error writing to client: %v", err)
					client.Close()
					delete(chatSsn.clients, client)
				}
			}
		}
		chatSsn.mutex.Unlock()
	}
}
