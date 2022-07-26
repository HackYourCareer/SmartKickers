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

type DispatcherMsgConverter struct {
	dispatcherMsg *dispatcherMsg
}

func (mg *DispatcherMsgConverter) Unpack(message []byte) error {
	return json.Unmarshal(message, mg.dispatcherMsg)
}

type DispatcherResConverter struct {
	dispatcherRes *dispatcherResponse
}

func (mg DispatcherResConverter) Pack(message []byte) ([]byte, error) {
	return json.Marshal(message)
}
