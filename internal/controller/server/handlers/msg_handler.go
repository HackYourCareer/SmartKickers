package handlers

import (
	"log"
	"net/http"
	"remote/internal/controller/server/adapter"
	"remote/internal/model"

	"github.com/gorilla/websocket"
)

func HandleTableMessages(game model.Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := Connect(w, r)
		if err != nil {
			log.Println(err)
		}

		defer c.Close()

		err = readTableMessage(c, game)
		if err != nil {
			log.Println(err)
		}
	}
}

func readTableMessage(c *websocket.Conn, game model.Game) error {
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
			log.Println("Connected")
			response, err := initialResponse(mt, mes.TableId)
			if err != nil {
				return err
			}
			c.WriteMessage(mt, response)
		}
		if mes.Goal != 0 {
			//TODO
			log.Println("GOOOOOL!")
			game.AddGoal(mes.Goal)
		}

	}
}

func initialResponse(connMsgType int, tableId string) ([]byte, error) {
	rec, err := adapter.PackGameId(tableId)
	return rec, err
}
