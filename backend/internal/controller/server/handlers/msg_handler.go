package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/controller/server/adapter"
	"github.com/HackYourCareer/SmartKickers/internal/model"

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

		var mes adapter.DispatcherMsg
		err = json.Unmarshal(message, &mes)
		if err != nil {
			return err
		}

		if mes.MsgType == "INITIAL" {
			response, err := json.Marshal(adapter.NewDispatcherResponse(mes.TableID))
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
