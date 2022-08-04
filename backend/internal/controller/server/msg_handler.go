package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

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

	message, err := adapter.Unpack(reader)
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

func (s server) ResetScoreHandler(w http.ResponseWriter, r *http.Request) {
	s.game.ResetScore()
}

//Incoming URL should be in the format:
// '/goal?action=[add/sub]&team=[1/2]'
// team 1 - White; team 2 - Blue
func (s server) ManipulateScoreHandler(w http.ResponseWriter, r *http.Request) {

	team := r.URL.Query().Get("team")

	teamID, err := strconv.Atoi(team)
	if err != nil {
		log.Fatal(err)
	}

	switch action := r.URL.Query().Get("action"); action {
	case "add":
		s.game.AddGoal(teamID)
	case "sub":
		// PLACEHOLDER
		//s.game.SubGoal(teamID)
	default:
		log.Panic("Wrong action")
	}
}
