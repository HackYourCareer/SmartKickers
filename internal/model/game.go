package model

type gameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

func New() *gameScore {
	return &gameScore{}
}

func (gameScore *GameScore) ResetScore() {
	gameScore.BlueScore = 0
	gameScore.WhiteScore = 0
}

func (gameScore *GameScore) AddGoal(teamID int) {
	switch teamID {
	case 1:
		gameScore.WhiteScore++
	case 2:
		gameScore.BlueScore++
	}
}
