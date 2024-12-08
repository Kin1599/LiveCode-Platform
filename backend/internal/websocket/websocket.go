package websocket

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000"
		return true
	},
}

var clients = make(map[*websocket.Conn]string)
var colors = make(map[string]string)
var broadcast = make(chan Message)
var mutex = &sync.Mutex{}

type Message struct {
	Type     string          `json:"type"`
	Text     string          `json:"text"`
	From     json.RawMessage `json:"from"`
	To       json.RawMessage `json:"to"`
	UserID   string          `json:"userId"`
	CursorX  int             `json:"cursorX"`
	CursorY  int             `json:"cursorY"`
	Color    string          `json:"color"`
	Nickname string          `json:"nickname"`
}

func generateColor() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("#%02X%02X%02X", r.Intn(256), r.Intn(256), r.Intn(256))
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
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

	mutex.Lock()
	clients[conn] = initialMessage.UserID
	colors[initialMessage.UserID] = generateColor()
	mutex.Unlock()

	for {
		var message Message
		if err := conn.ReadJSON(&message); err != nil {
			mutex.Lock()
			delete(clients, conn)
			delete(colors, clients[conn])
			mutex.Unlock()
			break
		}
		message.UserID = clients[conn]
		message.Color = colors[message.UserID]
		broadcast <- message
	}
}

func HandleMessages() {
	for {
		message := <-broadcast

		if message.Type == "cursor" && len(message.Text) > 100 {
			message.Text = message.Text[:100]
		}

		mutex.Lock()
		for client, userID := range clients {
			if userID != message.UserID {
				err := client.WriteJSON(message)
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
		mutex.Unlock()
	}
}
