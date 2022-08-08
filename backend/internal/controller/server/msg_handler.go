package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
	"github.com/gorilla/websocket"
)

const (
	messageTypeText = 1
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
			err = c.WriteMessage(messageTypeText, response)
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

func (s server) SendScoreHandler(w http.ResponseWriter, r *http.Request) {

	closeConnChan := make(chan bool)

	var upgrader websocket.Upgrader
	// TODO: We should check the origin in the future. For now we enable every connection to the server.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer c.Close()

	if err := s.writeScoreMessage(c); err != nil {
		log.Println(err)
	}

	go waitForError(c, closeConnChan)

	for {
		select {
		case <-s.game.GetChannel():
			if err := s.writeScoreMessage(c); err != nil {
				log.Println(err)
				break
			}
		case <-closeConnChan:
			return

		}
	}
}

func (s *server) writeScoreMessage(c *websocket.Conn) error {
	gameScoreMsg, err := json.Marshal(s.game.GetScore())
	if err != nil {
		return err
	}

	err = c.WriteMessage(messageTypeText, gameScoreMsg)

	return err
}

func waitForError(c *websocket.Conn, ch chan bool) {
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			ch <- true
			return
		}

	}
}
