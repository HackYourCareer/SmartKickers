package model

type Game struct {
	score gameScore
}

type gameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

func New() *Game {
	return &Game{}
}

func (gameScore *gameScore) ResetScore() {
	gameScore.BlueScore = 0
	gameScore.WhiteScore = 0
}

func (game *Game) AddGoal(teamID int) {
	switch teamID {
	case 1:
		game.score.WhiteScore++
	case 2:
		game.score.BlueScore++
	}
}
