package presentation

import (
	"golang.org/x/net/websocket"
)

type WebSocketServer struct {
	connections map[*websocket.Conn]bool
}

func NewWebsocketServer() *WebSocketServer {
	return &WebSocketServer{
		connections: make(map[*websocket.Conn]bool),
	}
}
