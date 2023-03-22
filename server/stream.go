package server

import (
	"github.com/gofiber/websocket/v2"
	"github.com/sourcegraph/jsonrpc2"
)

type wsObjectStream struct {
	conn *websocket.Conn
}

func (server *wsObjectStream) WriteObject(obj interface{}) error {
	return server.conn.WriteJSON(obj)
}

func (server *wsObjectStream) ReadObject(obj interface{}) error {
	return server.conn.ReadJSON(obj)
}

func (server *wsObjectStream) Close() error {
	return server.conn.Close()
}

func newWSObjectStream(conn *websocket.Conn) jsonrpc2.ObjectStream {
	return &wsObjectStream{conn}
}
