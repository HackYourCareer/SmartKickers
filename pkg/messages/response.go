package messages

import "encoding/json"

type DispatcherRes struct {
	GameId    int `json:"start,omitempty"`
	GameEnded int `json:"end,omitempty"`
}

func SendResponse(dr *DispatcherRes) {
	_, err := json.Marshal(DispatcherRes{GameId: dr.GameId, GameEnded: dr.GameEnded})
	if err != nil {
		panic(err)
	}
}
