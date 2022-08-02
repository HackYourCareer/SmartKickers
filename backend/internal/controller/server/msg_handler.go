package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
	"github.com/gorilla/websocket"
)

const (
	messageText = 1
)

func (s server) TableMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader websocket.Upgrader
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	defer c.Close()

	for {
		_, receivedMsg, err := c.NextReader()
		if err != nil {
			log.Println(err)
			continue
		}
		response, err := s.createResponse(receivedMsg)

		if err != nil {
			log.Println(err)
			continue
		}
		if response != nil {
			err = c.WriteMessage(messageText, response)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func (s server) createResponse(reader io.Reader) ([]byte, error) {

	message, err := adapter.Unpack(reader) //	Unpack will return our internal message type
	if err != nil {
		return nil, err
	}
	switch message.Category {
	case adapter.MsgInitial:
		return json.Marshal(adapter.NewDispatcherResponse(message.TableID))
	case adapter.MsgGoal:
		s.game.AddGoal(message.Team)
		return nil, nil
	default:
		return nil, BadMessageError{}
	}
}

type BadMessageError struct{}

func (err BadMessageError) Error() string {
	return "msg_handlers/createResponse(): Unrecognized message type"
}
