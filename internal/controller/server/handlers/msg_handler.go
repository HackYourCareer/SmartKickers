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

func readTableMessage(c *websocket.Conn) error {
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			return err
		}

		mes, err := adapter.Unpack(message)
		if err != nil {
			return err
		}

		if mes.MsgType == "INITIAL" {
			response, err := initialResponse(mt, mes.TableId)
			if err != nil {
				return err
			}
			c.WriteMessage(mt, response)
		}
		if mes.Goal != 0 {
			//TODO
			game.AddGoal(mes.Goal)
		}

	}
}

func initialResponse(connMsgType int, tableId string) ([]byte, error) {
	rec, err := adapter.PackGameId(tableId)
	return rec, err
}
