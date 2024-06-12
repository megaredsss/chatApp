package pkg

import "github.com/gorilla/websocket"

// Client need to add
type Client struct {
	ws     *websocket.Conn
	ch     chan *Message
	doneCh chan bool
}
