package model

type Game struct {
	score gameScore
}

type gameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

func (gS *gameScore) ResetScore() {
	gS.BlueScore = 0
	gS.WhiteScore = 0
}

func (g *Game) AddGoal(teamID int) {
	switch teamID {
	case 1:
		g.score.WhiteScore++
	case 2:
		g.score.BlueScore++
	}
}
