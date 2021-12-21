package websocket

import (
	gubrak "github.com/novalagung/gubrak/v2"
)

func ejectConnection(currentConn *WebSocketConnection) {
	filtered := gubrak.From(Connections).Reject(func(each *WebSocketConnection) bool {
		return each == currentConn
	}).Result()
	Connections = filtered.([]*WebSocketConnection)
}

func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
	//akan di broadcast ke semua users yang aktif
	for _, eachConn := range Connections {
		

		eachConn.WriteJSON(SocketResponse{
			From: currentConn.Email,
			Type: kind,
			Message: message,
		})
	}
}




