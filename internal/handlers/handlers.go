package handlers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *(http.Request)) bool { return true },
}

// Defines the response sent back from websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WebSocketConn struct {
	*websocket.Conn
}

// Payload for sendbing back Websocket Information to User
type WebSocketPayload struct {
	Action   string `json: "action"`
	Username string `json: "username"`
	Message  string `json: "message"`
	Conn     WebSocketConn
}

// upgrade standard responsewriter, request, response header to a websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to websocket server.")
	var response WsJsonResponse

	response.Message = "Connected!"

	err = ws.WriteJSON(response)

	if err != nil {
		log.Println(err)
	}

}
