package messages

type DispatcherMsg struct {
	MsgType   string  `json:"type"`
	Origin    string  `json:"origin"`
	TableId   int     `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Timestamp string  `json:"timestamp"`
	Goal      int     `json:"goal"`
	Sequence  string  `json:"Sequence"`
}
