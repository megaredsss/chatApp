package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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
		messageType, r, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			return
		}
		err = conn.WriteMessage(messageType, r)
		if err != nil {
			log.Println("write error: ", err)
			return
		}
	}
}
