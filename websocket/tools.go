package websocket

import (
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	gubrak "github.com/novalagung/gubrak/v2"
)

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection, params Params) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error websocket personal chat recover : ", r)
		}
	}()

	broadcastMessage(currentConn, IS_ONLINE, "")
	if err := params.UpdateStatus(currentConn, ONLINE); err != nil {
		log.Println("Error update status users ", ONLINE, " : ", err.Error())
	}

	for {
		payload := SocketPayload{}
		if err := currentConn.ReadJSON(&payload); err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				// update menjadi offline
				if err := params.UpdateStatus(currentConn, OFFLINE); err != nil {
					log.Println("Error update status user ", OFFLINE, " : ", err.Error())
				}
				broadcastMessage(currentConn, HAS_LEAVE, "")
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

		if payload.Command == PERSONAL_CHAT {
			personalMessage(currentConn, payload)
		} else if payload.Command == BROADCAST {
			broadcastMessage(currentConn, BROADCAST, payload.Message)
		}
	}
}

func ejectConnection(currentConn *WebSocketConnection) {
	filtered := gubrak.From(Connections).Reject(func(each *WebSocketConnection) bool {
		return each == currentConn
	}).Result()
	Connections = filtered.([]*WebSocketConnection)
}

func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
	//akan di broadcast ke semua users yang aktif
	for _, eachConn := range Connections {
		
		if eachConn == currentConn {
			// mencegah pengiriman ke diri sendiri
			continue
		}

		eachConn.WriteJSON(SocketResponse{
			From: currentConn.Email,
			Type: kind,
			Message: message,
		})
	}
}

func personalMessage(currentConn *WebSocketConnection, payload SocketPayload) {
	// ditujukan ke users tujuan
	for _, eachConn := range Connections {
		if eachConn.Email == payload.SendToEmail && eachConn.Role == payload.TargetRole {
			
			eachConn.WriteJSON(SocketResponse{
				From: currentConn.Email,
				Type: CHAT,
				Message: payload.Message,
			})
		}
	}
}

func updateDB(currentConn *WebSocketConnection, params Params, value string) error {
	query, args, err := sq.Update(params.Table).Set(params.Field, value).Where(sq.Eq{"email":[]string{currentConn.Email}}).ToSql()
	if err != nil {
		log.Println("Squirel error when update socket: ", err.Error())
		return err 
	}

	_, err = params.DB.Exec(query, args...)
	if err != nil {
		log.Println("An error occurred while updating websocket: ", err.Error())
		return err
	} 
	return nil
}




