package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"remote/internal/controller/server/adapter"
	"remote/internal/model"

	"github.com/gorilla/websocket"
)

func TableMessages(game model.Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := connectWebSocket(w, r)
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
			response, err := initialResponse(mes.TableID)
			if err != nil {
				return err
			}

			err = c.WriteMessage(mt, response)

			if err != nil {
				return err
			}

		}
		if mes.Goal != 0 {
			game.AddGoal(mes.Goal)
		}

	}
}

func initialResponse(tableID string) ([]byte, error) {
	rec, err := json.Marshal(adapter.NewDisRes(tableID))
	return rec, err
}
