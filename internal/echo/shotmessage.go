package internal

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"remote/pkg/messages"
)

func ReceiveShotMsg(w http.ResponseWriter, r *http.Request) {

	e, c := connect(w, r)
	if e != nil {
		return
	}

	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		_ = handleMessages(message)
	}
}

func connect(w http.ResponseWriter, r *http.Request) (error, *websocket.Conn) {
	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return err, nil
	}
	return err, c
}

func handleMessages(message []byte) messages.ShotParams {
	var shotMsg messages.ShotMsg
	_ = json.Unmarshal(message, &shotMsg)
	var shotParams messages.ShotParams
	_ = json.Unmarshal(shotMsg.Params[0], &shotParams)

	log.Printf("Shot with speeed of %v\n", shotParams.Speed)
	return shotParams
}
