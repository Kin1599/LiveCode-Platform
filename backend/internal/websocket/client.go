package websocket

import (
	"encoding/json"
	"fmt"
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
	clients   map[*websocket.Conn]string
	colors    map[string]string
	broadcast chan Message
	mutex     sync.Mutex
	history   []Message
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
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
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
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
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
		http.Error(w, "session_id is required", http.StatusBadRequest)
		fmt.Println("session_id is required")
		return
	}

	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		http.Error(w, "Invalid session_id value", http.StatusBadRequest)
		return
	}

	chatSession, ok := chatSessions[sessionID]
	if !ok {
		chatSession = &ChatSession{
			clients:   make(map[*websocket.Conn]string),
			colors:    make(map[string]string),
			broadcast: make(chan Message),
			history:   []Message{},
		}
		chatSessions[sessionID] = chatSession

		go chatSession.HandleMessages()
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	var initialMessage ChatMessage
	if err := conn.ReadJSON(&initialMessage); err != nil {
		fmt.Println("Error reading initial message:", err)
		return
	}

	chatSession.mutex.Lock()
	chatSession.clients[conn] = initialMessage.UserID
	chatSession.colors[initialMessage.UserID] = generateColor()

	historyMessage := Message{
		Type:    "history",
		History: chatSession.history,
	}

	if err := conn.WriteJSON(historyMessage); err != nil {
		fmt.Println("Error sending history message: ", err)
		return
	}

	chatSession.mutex.Unlock()

	client := newClient(conn)

	go client.writePump()
	go client.readPump()

	register <- client
}

func (chatSsn *ChatSession) HandleMessages() {
	message := <-chatSsn.broadcast

	chatSsn.mutex.Lock()
	for client, userID := range chatSsn.clients {
		if userID != message.UserID {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(chatSsn.clients, client)
			}
		}
	}
	chatSsn.mutex.Unlock()
}
