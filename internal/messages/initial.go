package messages

type InitialMessage struct {
	MsgType string `json:"type"`
	Origin string `json:"origin"`
	TableId int `json:"id"` 
}