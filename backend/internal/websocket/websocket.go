package websocket

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000"
		return true
	},
}

type Session struct {
	clients   map[*websocket.Conn]string
	colors    map[string]string
	broadcast chan Message
	mutex     sync.Mutex
	history   []Message
}

var sessions = make(map[uuid.UUID]*Session)

type Message struct {
	Type      string          `json:"type"`
	Text      string          `json:"text"`
	From      json.RawMessage `json:"from"`
	To        json.RawMessage `json:"to"`
	UserID    string          `json:"userId"`
	CursorX   int             `json:"cursorX"`
	CursorY   int             `json:"cursorY"`
	Color     string          `json:"color"`
	Nickname  string          `json:"nickname"`
	SessionID string          `json:"sessionId"`
	History   []Message       `json:"history"`
}

func generateColor() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("#%02X%02X%02X", r.Intn(256), r.Intn(256), r.Intn(256))
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
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

	session, ok := sessions[sessionID]
	if !ok {
		session = &Session{
			clients:   make(map[*websocket.Conn]string),
			colors:    make(map[string]string),
			broadcast: make(chan Message),
			history:   []Message{},
		}
		sessions[sessionID] = session

		go session.HandleMessages()
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}

	defer conn.Close()

	var initialMessage Message
	if err := conn.ReadJSON(&initialMessage); err != nil {
		fmt.Println("Error reading initial message:", err)
		return
	}

	session.mutex.Lock()
	session.clients[conn] = initialMessage.UserID
	session.colors[initialMessage.UserID] = generateColor()

	historyMessage := Message{
		Type:    "history",
		History: session.history,
	}

	if err := conn.WriteJSON(historyMessage); err != nil {
		fmt.Println("Error sending history message: ", err)
		return
	}

	session.mutex.Unlock()

	for {
		var message Message
		if err := conn.ReadJSON(&message); err != nil {
			session.mutex.Lock()
			userID := session.clients[conn]
			delete(session.colors, userID)
			delete(session.clients, conn)
			session.broadcast <- Message{
				Type:   "removeCursor",
				UserID: session.clients[conn],
			}
			session.mutex.Unlock()
			break
		}
		message.UserID = session.clients[conn]
		message.Color = session.colors[message.UserID]
		session.broadcast <- message
		session.history = append(session.history, message)
	}
}

func (s *Session) HandleMessages() {
	for {
		message := <-s.broadcast

		if message.Type == "cursor" && len(message.Text) > 100 {
			message.Text = message.Text[:100]
		}

		s.mutex.Lock()
		for client, userID := range s.clients {
			if userID != message.UserID {
				err := client.WriteJSON(message)
				if err != nil {
					client.Close()
					delete(s.clients, client)
				}
			}
		}
		s.mutex.Unlock()
	}
}
