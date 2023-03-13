package model

import "github.com/gorilla/websocket"

// Connection represents an active connection and all data associated with it
type Connection struct {
	WebSockerConn     *websocket.Conn
	SubscribedChannel chan []uint8
}
