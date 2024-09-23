package pkg

import "github.com/gorilla/websocket"

// Client need to add
type Client struct {
	ID string
	Ws *websocket.Conn
}
