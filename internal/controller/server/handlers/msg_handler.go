package handlers

import (
	"log"
	"net/http"
	"remote/internal/controller/server/adapter"

	"github.com/gorilla/websocket"
)

func HandleTableMessages(w http.ResponseWriter, r *http.Request) {
	c, err := Connect(w, r)
	if err != nil {
		log.Println(err)
	}

	defer c.Close()

	readTableMessage(c)
}

func readTableMessage(c *websocket.Conn) {
	for {
		connMsgType, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		//mes := adapter(message)
		log.Print(message)
		mes := adapter.DispatcherMsg{}

		checkMessageType(connMsgType, mes)
	}
}

func checkMessageType(connMsgType int, msg adapter.DispatcherMsg) {
	if msg.MsgType == "INITIAL" {
		initialResponse(connMsgType, msg)
	} else if msg.Goal != 0 {
		goalResponse(msg.Goal)
	}
}

func initialResponse(connMsgType int, msg adapter.DispatcherMsg) {

}

func goalResponse(team int) {

}
