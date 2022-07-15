package messages

type DispatcherResMsg struct {
	GameId    int `json:"start"`
	GameEnded int `json:"end"`
}
