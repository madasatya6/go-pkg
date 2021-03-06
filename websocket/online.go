package websocket

import (
	"log"
	"strings"
)


func handleOnline(currentConn *WebSocketConnection, connections []*WebSocketConnection, params Params) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error websocket recover : ", r)
		}
	}()

	broadcastMessage(currentConn, IS_ONLINE, "")
	updateDB(currentConn, params, ONLINE)

	for {
		payload := SocketPayload{}
		if err := currentConn.ReadJSON(&payload); err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, HAS_LEAVE, "")
				updateDB(currentConn, params, OFFLINE)
				ejectConnection(currentConn)
				return 
			}

			log.Println("Error Read JSON websocket: ", err.Error())
			continue
		}
	}
}


