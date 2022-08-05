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
	messageText     = 1
	attributeTeam   = "team"
	attributeAction = "action"
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

// ManipulateScoreHandler is a handler for manipulation of the score.
// Incoming URL should be in the format: '/goal?action=[add/sub]&team=[1/2]'.
// Team ID 1 stands for white and 2 for blue.
func (s server) ManipulateScoreHandler(w http.ResponseWriter, r *http.Request) {

	team := r.URL.Query().Get(attributeTeam)

	teamID, err := strconv.Atoi(team)
	if err != nil || !isValidTeamID(teamID) {
		handleError(w, http.StatusBadRequest, "Team ID has to be a number either 1 or 2")
		return
	}

	switch action := r.URL.Query().Get(attributeAction); action {
	case "add":
		s.game.AddGoal(teamID)
	case "sub":
		// TODO: s.game.SubGoal(teamID)
	default:
		handleError(w, http.StatusBadRequest, "Wrong action")
	}
}

func handleError(w http.ResponseWriter, header int, msg string) {
	log.Println("Error handling request: ", msg)
	w.WriteHeader(header)
	w.Write([]byte(msg))
}

func isValidTeamID(teamID int) bool {
	return (teamID == 1 || teamID == 2)
}
