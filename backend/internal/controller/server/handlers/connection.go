package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func Connect(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var upgrader websocket.Upgrader
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } //	Allow all connections by default
	c, err := upgrader.Upgrade(w, r, nil)

	return c, err
}
