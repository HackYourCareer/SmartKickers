package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const messageTypeText = 1

func (s server) TableMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader websocket.Upgrader
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
	}

	defer c.Close()

	for {
		_, receivedMsg, err := c.NextReader()
		if err != nil {
			log.Error(err)
			continue
		}
		response, err := s.createResponse(receivedMsg)

		if err != nil {
			log.Debug(err)
			continue
		}
		if response != nil {
			err = c.WriteMessage(messageTypeText, response)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}

func (s server) createResponse(reader io.Reader) ([]byte, error) {
	message, err := adapter.UnpackDispatcherMsg(reader)
	if err != nil {
		return nil, err
	}
	switch message.Category {
	case adapter.MsgInitial:
		return json.Marshal(adapter.NewDispatcherResponse(message.TableID))
	case adapter.MsgGoal:
		err := s.game.AddGoal(message.Team)
		return nil, err
	default:
		return nil, fmt.Errorf("unrecognized message type %d", message.Category)
	}
}

func (s server) ShotParametersHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader websocket.Upgrader
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
	}

	defer c.Close()

	for {
		_, receivedMsg, err := c.NextReader()
		if err != nil {
			log.Error(err)
			continue
		}

		shot, err := adapter.UnpackShotMsg(receivedMsg)
		if err != nil {
			log.Error(err)
		}
		if s.game.IsFastestShot(shot.Speed) {
			s.game.SaveFastestShot(shot)
		}
		log.Debug(shot)
	}
}
