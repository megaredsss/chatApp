package pkg

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsCreater(response http.ResponseWriter, request *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(response, request, nil)
	if err != nil {
		http.NotFound(response, request)
		return
	}
	fmt.Println(conn)
}
