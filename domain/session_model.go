package domain

import "github.com/gorilla/websocket"

type SessionSocket struct {
	Sid string
	Ws *websocket.Conn
}