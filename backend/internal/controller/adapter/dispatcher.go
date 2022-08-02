package adapter

import (
	"encoding/json"
	"io"
)

type Category int

const (
	None Category = iota
	Initial
	Goal
	Position
)

type Message struct {
	Category Category
	TableID  string
	Team     int
}

type dispatcherMsg struct {
	MsgType   string  `json:"type,omitempty"`
	Origin    string  `json:"origin,omitempty"`
	TableID   string  `json:"id,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
	Goal      int     `json:"goal,omitempty"`
	Sequence  string  `json:"Sequence,omitempty"`
}

type initialResponse struct {
	GameID    string `json:"start,omitempty"`
	GameEnded int    `json:"end,omitempty"`
}

func Unpack(message io.Reader) (*Message, error) {
	var tableMessage dispatcherMsg

	err := json.NewDecoder(message).Decode(&tableMessage)
	if err != nil {
		return &Message{}, err
	}

	return &Message{
		Category: tableMessage.getMessageCategory(),
		TableID:  tableMessage.TableID,
		Team:     tableMessage.Goal,
	}, nil
}

func (dispMsg dispatcherMsg) getMessageCategory() Category {
	if dispMsg.MsgType == "INITIAL" {
		return Initial
	}
	if dispMsg.Goal != 0 {
		return Goal
	}
	return None
}

func NewDispatcherResponse(tableID string) *initialResponse {
	dr := new(initialResponse)
	dr.GameID = tableID
	return dr
}
