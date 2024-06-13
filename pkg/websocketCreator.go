package pkg

import (
	models "chatApp/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	models.Room
}
type client struct {
	models.Client
}
type message struct {
	models.Message
}

// Go my favorite language
// var client Client
// var message Message
// var room Room

// HTTP -> WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

func startWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in startWs", err)
		return
	}
	for {
		// Reading
		messageType, r, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			return
		}
		// Writing
		err = conn.WriteMessage(messageType, r)
		if err != nil {
			log.Println("write error: ", err)
			return
		}
	}
}

func (room room) Start() {
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
			fmt.Println("Size of Connection room: ", len(room.Clients))
			for client, _ := range room.Clients {
				fmt.Println(client)
				client.Ws.WriteJSON(models.Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-room.Unregister:
			delete(room.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(room.Clients))
			for client, _ := range room.Clients {
				client.Ws.WriteJSON(models.Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-room.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range room.Clients {
				if err := client.Ws.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
