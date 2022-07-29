package server

import (
	"io"
	"log"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/controller/server/adapter"
	"github.com/gorilla/websocket"
)

func (s server) TableMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader websocket.Upgrader
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err) //TODO change logging to logrus
	}

	defer c.Close()

	for {
		_, receivedMsg, err := c.NextReader()
		if err != nil {
			log.Println(err) //TODO change logging to logrus
			continue
		}
		err = s.readAndRespond(receivedMsg)
		if err != nil {
			log.Println(err) //TODO change logging to logrus
			continue
		}
	}
}

func (s server) readAndRespond(r io.Reader) error {

	message, err := adapter.Unpack(r) //	Unpack will return our internal message type
	if err != nil {
		return err //TODO change logging to logrus
	}

	switch message.Category {
	case adapter.Initial:
		initial(message.TableID)
	case adapter.Goal:
		s.game.AddGoal(message.Team)
	default:
		log.Println("TableMessagesHandler: Bad message") //TODO change logging to logrus
	}

	return nil
}

func initial(tableID string) {
	log.Println("initial")
}

/*func (s server) TableMessages(w http.ResponseWriter, r *http.Request) {
	c, err := connectWebSocket(w, r)
	if err != nil {
		log.Println(err)
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		var mes adapter.DispatcherMsg
		err = json.Unmarshal(message, &mes)
		if err != nil {
			log.Println(err)
		}

		if mes.MsgType == "INITIAL" {
			response, err := json.Marshal(adapter.NewDispatcherResponse(mes.TableID))
			if err != nil {
				log.Println(err)
			}

			err = c.WriteMessage(mt, response)

			if err != nil {
				log.Println(err)
			}

		}
		if mes.Goal != 0 {
			s.game.AddGoal(mes.Goal)
		}

	}
}*/
