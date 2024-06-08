package pkg

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsCreater(response http.ResponseWriter, request *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(response, request, nil)
	if err != nil {
		http.NotFound(response, request)
		return
	}
	fmt.Println(conn)
}
