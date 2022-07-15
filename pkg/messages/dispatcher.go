package messages

type DispatcherReq struct {
	MsgType string  `json:"type,omitempty"`
	Origin  string  `json:"origin,omitempty"`
	TableId int     `json:"id,omitempty"`
	X       float64 `json:"x,omitempty"`
	Y       float64 `json:"y,omitempty"`
	Goal    int     `json:"goal,omitempty"`
}
