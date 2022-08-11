package adapter

import (
	"encoding/json"
	"io"
)

type ShotMessage struct {
	Speed float64
	Team  int
}

type tableShotParams struct {
	TimeStart int     `json:"TimeStart,omitempty"`
	TimeEnd   int     `json:"TimeEnd,omitempty"`
	Speed     float64 `json:"Speed,omitempty"`
	StartArea int     `json:"StartArea,omitempty"`
	EndArea   int     `json:"EndArea,omitempty"`
	GameID    string  `json:"GameID,omitempty"`
	Sequence  int     `json:"Sequence,omitempty"`
	ID        string  `json:"ID,omitempty"`
}

type tableShotMsg struct {
	Mode        string            `json:"mode,omitempty"`
	MessageType string            `json:"messageType,omitempty"`
	Params      []json.RawMessage `json:"messages,omitempty"`
}

func UnpackShotMsg(message io.Reader) (ShotMessage, error) {
	var (
		shotMessage tableShotMsg
		params      tableShotParams
	)

	err := json.NewDecoder(message).Decode(&shotMessage)
	if err != nil {
		return ShotMessage{}, err
	}

	err = json.Unmarshal(shotMessage.Params[0], &params)
	if err != nil {
		return ShotMessage{}, err
	}

	return ShotMessage{
		Speed: params.Speed,
		Team:  decodeTeam(params.StartArea),
	}, nil
}

func decodeTeam(a int) int {
	return a%2 + 1
}
