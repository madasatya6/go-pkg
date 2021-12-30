package websocket

import (
	"log"
	"strings"
)


func handlePersonalChat(currentConn *WebSocketConnection, connections []*WebSocketConnection, params Params) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error websocket personal chat recover : ", r)
		}
	}()

	for {
		payload := SocketPayload{}
		if err := currentConn.ReadJSON(&payload); err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				ejectConnection(currentConn)
				return 
			}

			log.Println("Error Read JSON websocket: ", err.Error())
			continue
		}

		// insert db chat
		if err := params.InsertChat(currentConn, payload); err != nil {
			log.Println("Error insert chat : ", err.Error())
		}
		personalMessage(currentConn, payload)
	}
}


