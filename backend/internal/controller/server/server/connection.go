package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func connectWebSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var upgrader websocket.Upgrader
	return upgrader.Upgrade(w, r, nil)
}
