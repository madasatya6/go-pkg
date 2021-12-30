package websocket

import (
	"log"
	"context"
	"net/http"
	"database/sql"

	"github.com/gorilla/websocket"
)

type M map[string]interface{}

const IS_ONLINE = "Online"
const HAS_LEAVE = "Leave"
const IS_JOIN_CHAT = "Join Chat" 
const IS_CHAT = "Chat"

var Connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message 	string 
	SendToEmail string // to email 
	TargetRole	string // admin & customer
	Command 	string // broadcast & personal
}

type SocketResponse struct {
	From 	string 
	Type 	string
	Message string 
}

type WebSocketConnection struct {
	*websocket.Conn 
	Email string 
	Role string
}

type Params struct {
	Feature string 
	DB 		*sql.DB
	Table 	string 
	Field 	string
	Context context.Context
	
	InsertChat func(currentConn *WebSocketConnection, payload SocketPayload) error
	UpdateStatus func(currentConn *WebSocketConnection, status string) error
}

// main function
func Connect(w http.ResponseWriter, r *http.Request, params Params) {
	currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		log.Println("Error init websocket : ", err.Error())
		return 
	}

	email := r.URL.Query().Get("email")
	role := r.URL.Query().Get("role")

	currentConn := WebSocketConnection{Conn: currentGorillaConn, Email: email, Role: role}
	Connections = append(Connections, &currentConn)

	go handleIO(&currentConn, Connections, params)
} 




