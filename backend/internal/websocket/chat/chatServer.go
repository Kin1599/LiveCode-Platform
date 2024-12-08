package chat

// type WsServer struct {
// 	clients    map[*Client]bool
// 	register   chan *Client
// 	unregister chan *Client
// 	broadcast  chan []byte
// }

var clients = make(map[*Client]bool)
var register = make(chan *Client)
var unregister = make(chan *Client)

var broadcast = make(chan []byte)

// // NewWebsocketServer creates a new WsServer type
// func NewWebsocketServer() *WsServer {
// 	return &WsServer{
// 		clients:    make(map[*Client]bool),
// 		register:   make(chan *Client),
// 		unregister: make(chan *Client),
// 		broadcast:  make(chan []byte),
// 	}
// }

// Run our websocket server, accepting various requests
func Run() {
	for {
		select {

		case client := <-register:
			registerClient(client)

		case client := <-unregister:
			unregisterClient(client)

		case message := <-broadcast:
			broadcastToClients(message)
		}

	}
}

func registerClient(client *Client) {
	clients[client] = true
}

func unregisterClient(client *Client) {
	if _, ok := clients[client]; ok {
		delete(clients, client)
	}
}

func broadcastToClients(message []byte) {
	for client := range clients {
		client.send <- message
	}
}
