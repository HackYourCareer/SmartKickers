package adapter

type DispatcherMsg struct {
	MsgType   string  `json:"type,omitempty"`
	Origin    string  `json:"origin,omitempty"`
	TableID   string  `json:"id,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
	Goal      int     `json:"goal,omitempty"`
	Sequence  string  `json:"Sequence,omitempty"`
}

type DispatcherResponse struct {
	GameID    string `json:"start,omitempty"`
	GameEnded int    `json:"end,omitempty"`
}

func NewDispatcherResponse(tableID string) *DispatcherResponse {
	dr := new(DispatcherResponse)
	dr.GameID = tableID
	return dr
}
