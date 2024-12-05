package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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
var broadcast = make(chan Message)
var mutex = &sync.Mutex{}

type Message struct {
	Type   string          `json:"type"`
	Text   string          `json:"text"`
	From   json.RawMessage `json:"from"`
	To     json.RawMessage `json:"to"`
	UserID string          `json:"userId"`
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
	mutex.Unlock()

	for {
		var message Message
		if err := conn.ReadJSON(&message); err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		message.UserID = clients[conn]
		broadcast <- message
	}
}

func HandleMessages() {
	for {
		message := <-broadcast
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
