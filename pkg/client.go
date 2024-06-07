package pkg

import "github.com/gorilla/websocket"

type Client struct {
	ID     int
	ws     *websocket.Conn
	ch     chan *Message
	doneCh chan bool
}
