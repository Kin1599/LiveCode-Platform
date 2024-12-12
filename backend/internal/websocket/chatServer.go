package websocket


var clients = make(map[*Client]bool)
var register = make(chan *Client)
var unregister = make(chan *Client)
var broadcast = make(chan []byte)


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
	delete(clients, client)
}

func broadcastToClients(message []byte) {
	for client := range clients {
		client.send <- message
	}
}