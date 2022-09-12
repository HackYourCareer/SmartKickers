package adapter

import (
	"encoding/json"
	"io"
)

type MsgCategory int

const (
	MsgNone MsgCategory = iota
	MsgInitial
	MsgGoal
	MsgPosition
)

type Message struct {
	Category MsgCategory
	TableID  string
	Team     int
	X        float64
	Y        float64
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

func UnpackDispatcherMsg(message io.Reader) (Message, error) {
	var tableMessage dispatcherMsg

	err := json.NewDecoder(message).Decode(&tableMessage)
	if err != nil {
		return Message{}, err
	}

	return Message{
		Category: tableMessage.getMessageCategory(),
		TableID:  tableMessage.TableID,
		Team:     tableMessage.Goal,
		X:        tableMessage.X,
		Y:        tableMessage.Y,
	}, nil
}

func (dispMsg dispatcherMsg) getMessageCategory() MsgCategory {
	if dispMsg.MsgType == "INITIAL" {
		return MsgInitial
	}

	if dispMsg.Goal != 0 {
		return MsgGoal
	}

	if dispMsg.MsgType == "" {
		return MsgPosition
	}
	return MsgNone
}

func NewDispatcherResponse(tableID string) *initialResponse {
	dr := new(initialResponse)
	dr.GameID = tableID

	return dr
}
