package model

import "github.com/gorilla/websocket"

type ConnectUser struct {
	Websocket *websocket.Conn
	ClientIP  string
}

func NewConnectUser(ws *websocket.Conn, clientIP string) *ConnectUser {
	return &ConnectUser{
		Websocket: ws,
		ClientIP:  clientIP,
	}
}
