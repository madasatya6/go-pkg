package websocket

import (
	"log"
	sq "github.com/Masterminds/squirrel"
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




