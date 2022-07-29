package adapter

import (
	"encoding/json"
	"io"
)

type Category int

const (
	Initial Category = iota
	Goal
	Position
)

type Message struct {
	Category Category
	TableID  string
	Team     int
}

type DispatcherMsg struct {
	MsgType   string  `json:"type,omitempty"`
	Origin    string  `json:"origin,omitempty"`
	TableID   string  `json:"id,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
	Goal      int     `json:"goal,omitempty"`
	Sequence  string  `json:"Sequence,omitempty"`
}

type DispatcherResponse struct {
	GameID    string `json:"start,omitempty"`
	GameEnded int    `json:"end,omitempty"`
}

func Unpack(message io.Reader) (*Message, error) {
	var tableMessage DispatcherMsg
	var newMessage Message

	err := json.NewDecoder(message).Decode(&tableMessage)
	if err != nil {
		return &newMessage, err
	}

	var cat Category
	if tableMessage.MsgType == "INITIAL" {
		cat = Initial
	} else if tableMessage.Goal != 0 {
		cat = Goal
	}

	newMessage.Category = cat
	newMessage.TableID = tableMessage.TableID
	newMessage.Team = tableMessage.Goal

	return &newMessage, nil
}

func NewDispatcherResponse(tableID string) *DispatcherResponse {
	dr := new(DispatcherResponse)
	dr.GameID = tableID
	return dr
}
