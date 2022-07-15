package messages

type ShotMsg struct {
	TimeStart int     `json:"TimeStart"`
	TimeEnd   int     `json:"TimeEnd"`
	Speed     float64 `json:"Speed"`
	StartArea int     `json:"StartArea"`
	GameID    int     `json:"GameID"`
	Sequence  int     `json:"Sequence"`
	ID        string  `json:"ID"`
}
