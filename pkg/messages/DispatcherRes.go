package messages

import "encoding/json"

type DispatcherResponse struct {
	GameId    string `json:"start,omitempty"`
	GameEnded int    `json:"end,omitempty"`
}

func SendResponse(dr *DispatcherResponse) []byte {
	res, err := json.Marshal(DispatcherResponse{GameId: dr.GameId, GameEnded: dr.GameEnded})
	if err != nil {
		panic(err)
	}
	return res
}

func (dr *DispatcherResponse) New(gameId string, endId int) {
	dr.GameId = gameId
	dr.GameEnded = endId
}
