package pkg

import "github.com/gorilla/websocket"

type Client struct {
	ws     *websocket.Conn
	ch     chan *Message
	doneCh chan bool
}
