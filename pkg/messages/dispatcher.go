package messages

type DispatcherReqMsg struct {
	MsgType string `json:"type"`
	Origin string `json:"origin"`
	TableId int `json:"id"`
	X float64 `json:"x"`
	Y float64`json:"y"`
	Goal int `json:"goal"`
}

type DispatcherResMsg struct {
	GameId int `json:"start"`
	GameEnded int `json:"end"`
}