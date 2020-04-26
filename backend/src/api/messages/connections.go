package messages

import (
	"github.com/gorilla/websocket"
)

var chatId2Conns = map[uint]map[*websocket.Conn]bool{}

func AddConnection(chatId uint, conn *websocket.Conn) {
	conns, chatHasConns := chatId2Conns[chatId]
	if !chatHasConns {
		chatId2Conns[chatId] = make(map[*websocket.Conn]bool)
		conns = chatId2Conns[chatId]
	}

	_, connectionFound := conns[conn]
	if !connectionFound {
		conns[conn] = true
	}
}

func RemoveConnection(conn *websocket.Conn) {
	for _, conns := range chatId2Conns {
		_, hasConnection := conns[conn]
		if hasConnection {
			delete(conns, conn)
		}
	}
}

func GetConnections(chatId uint) []*websocket.Conn {
	var connections []*websocket.Conn
	for conn := range chatId2Conns[chatId] {
		connections = append(connections, conn)
	}

	return connections
}
