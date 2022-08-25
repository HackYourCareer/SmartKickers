package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
	"github.com/HackYourCareer/SmartKickers/internal/model"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	messageTypeText = 1
	attributeTeam   = "team"
	attributeAction = "action"
)

func (s server) TableMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader websocket.Upgrader

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer c.Close()

	for {
		_, receivedMsg, err := c.NextReader()
		if err != nil {
			log.Error(err)

			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				log.Error("Closing TableMessagesHandler")
				return
			}

			continue
		}

		response, err := s.createResponse(receivedMsg)

		if err != nil {
			log.Error(err)
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

func (s server) SendScoreHandler(w http.ResponseWriter, r *http.Request) {
	closeConnChan := make(chan error)

	var upgrader websocket.Upgrader
	// TODO: We should check the origin in the future. For now we enable every connection to the server.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer c.Close()

	if err := c.WriteJSON(s.game.GetScore()); err != nil {
		log.Error(err)
	}

	go waitForError(c, closeConnChan)

	for {
		select {
		case score := <-s.game.GetScoreChannel():
			if err := c.WriteJSON(score); err != nil {
				log.Error(err)
			}
		case err := <-closeConnChan:
			log.Error(err)
			return
		}
	}
}

func waitForError(c *websocket.Conn, ch chan error) {
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			ch <- err
			return
		}
	}
}

// ManipulateScoreHandler is a handler for manipulation of the score.
// Incoming URL should be in the format: '/goal?action=[add/sub]&team=[1/2]'.
// Team ID 1 stands for white and 2 for blue.
func (s server) ManipulateScoreHandler(w http.ResponseWriter, r *http.Request) {
	team := r.URL.Query().Get(attributeTeam)

	teamID, err := strconv.Atoi(team)
	if err != nil || !isValidTeamID(teamID) {
		err := writeHTTPError(w, http.StatusBadRequest, "Team ID has to be a number either 1 or 2")
		if err != nil {
			log.Error(err)
		}

		return
	}

	switch action := r.URL.Query().Get(attributeAction); action {
	case "add":
		err := s.game.AddGoal(teamID)
		if err != nil {
			log.Error(err)
		}
	case "sub":
		err := s.game.SubGoal(teamID)
		if err != nil {
			log.Error(err)
		}
	default:
		err := writeHTTPError(w, http.StatusBadRequest, "Wrong action")
		if err != nil {
			log.Error(err)
		}
	}
}

func writeHTTPError(w http.ResponseWriter, header int, msg string) error {
	log.Error("Error handling request: ", msg)
	w.WriteHeader(header)
	_, err := w.Write([]byte(msg))

	return err
}

func isValidTeamID(teamID int) bool {
	return (teamID == model.TeamWhite || teamID == model.TeamBlue)
}
