package adapter

import (
	"encoding/json"
)

type dispatcherMsg struct {
	MsgType   string  `json:"type,omitempty"`
	Origin    string  `json:"origin,omitempty"`
	TableId   string  `json:"id,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
	Goal      int     `json:"goal,omitempty"`
	Sequence  string  `json:"Sequence,omitempty"`
}

type dispatcherResponse struct {
	GameId    string `json:"start,omitempty"`
	GameEnded int    `json:"end,omitempty"`
}

func Unpack(message []byte) (*dispatcherMsg, error) {
	mg := new(dispatcherMsg)
	return mg, json.Unmarshal(message, &mg)
}

func NewDisRes(tableId string) *dispatcherResponse {
	dr := new(dispatcherResponse)
	dr.GameId = tableId
	return dr
}

func PackGameId(tableId string) ([]byte, error) {
	return json.Marshal(NewDisRes(tableId))
}
