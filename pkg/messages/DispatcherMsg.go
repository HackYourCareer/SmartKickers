package messages

type DispatcherMsg struct {
	TypeMess  string `json:"type"`
	Origin    string `json:"origin"`
	Id        string `json:"id"`
	XVal      string `json:"X"`
	YVal      string `json:"Y"`
	Timestamp string `json:"timestamp"`
	GameID    string `json:"GameID"`
	Sequence  string `json:"Sequence"`
}
