package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// channel to track websocket payloads
var wsChan = make(chan WebSocketPayload)

// track connected clients
var clients = make(map[WebSocketConnection]string)

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

type WebSocketConnection struct {
	*websocket.Conn
}

// Payload for sendbing back Websocket Information to User
type WebSocketPayload struct {
	Action   string `json: "action"`
	Username string `json: "username"`
	Message  string `json: "message"`
	Conn     WebSocketConnection
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

	// client successfully connected and recieved response
	// add them to list of connections and handle it there
	clientConnection := WebSocketConnection{
		Conn: ws,
	}

	clients[clientConnection] = ""

	// starts go routine that listen to incoming payloads
	go ListenForWS(&clientConnection)

	if err != nil {
		log.Println(err)
	}
}

// Listen for WebSocket Connections
func ListenForWS(conn *WebSocketConnection) {
	// logs error when the function stops and recovers
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()
	fmt.Println("Listening for websocket connection. Current clients ", clients)

	// payload that conforms to our custom WS payload
	var payload WebSocketPayload

	// infinite for loop to listen
	for {
		err := conn.ReadJSON(&payload)

		if err != nil {

		} else {
			// add connection pointer to the WebSocket Payload
			payload.Conn = *conn

			// send payload to our websocket channel
			wsChan <- payload
		}
	}
}

// Listen to the WebSocket CHANNEL
func ListenForWSChannel() {
	var response WsJsonResponse

	for {
		// storing websocket payload coming from wsChan
		event := <-wsChan

		// responds events sent to the channel to all users

		response.Action = "event"
		response.Message = fmt.Sprintf("Message received, action was %s. Message was: %s", event.Action, event.Message)

		// broadcast to all users
		broadcastToAll(response)
	}

}

// Broadcast to all users
func broadcastToAll(message WsJsonResponse) {

	// loop through all connected clients and broadcast to them
	for clientWS := range clients {

		err := clientWS.WriteJSON(message)

		// handle if client errored / disconnected
		if err != nil {
			fmt.Println("Websocket errored")

			// close their WS connection
			_ = clientWS.Close()

			// remove the client that errored
			delete(clients, clientWS)
		}
	}

}
