package messages

type ShotParams struct {
	TimeStart int     `json:"TimeStart,omitempty"`
	TimeEnd   int     `json:"TimeEnd,omitempty"`
	Speed     float64 `json:"Speed,omitempty"`
	StartArea int     `json:"StartArea,omitempty"`
	GameID    int     `json:"GameID,omitempty"`
	Sequence  int     `json:"Sequence,omitempty"`
	ID        string  `json:"ID,omitempty"`
}

type ShotMsg struct {
	Mode        string   `json:"mode,omitempty"`
	MessageType string   `json:"messageType,omitempty"`
	Params      []string `json:"messages,omitempty"`
}
