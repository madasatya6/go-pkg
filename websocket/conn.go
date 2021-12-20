package websocket

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

const JUST_ONLINE = "Online"
const JUST_LEAVE = "Leave"
const JUST_JOIN_CHAT = "Join Chat" 
const MESSAGE_CHAT = "Chat"

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string 
}

type SocketResponse struct {
	From 	string 
	Type 	string
	Message string 
}

type WebSocketConnection struct {
	*websocket.Conn 
	Username string 
	Admin bool
}


