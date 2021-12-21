package websocket

import (
	"log"
	"strings"
)


func handleOnline(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error websocket recover : ", r)
		}
	}()

	broadcastMessage(currentConn, IS_ONLINE, "")

	for {
		payload := SocketPayload{}
		if err := currentConn.ReadJSON(&payload); err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, HAS_LEAVE, "")
				ejectConnection(currentConn)
				return 
			}

			log.Println("Error Read JSON websocket: ", err.Error())
			continue
		}
	}
}


