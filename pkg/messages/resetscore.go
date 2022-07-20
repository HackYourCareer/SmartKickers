package messages

func (gScore *GameScore) ResetScore() {
	gScore.WhiteScore = 0
	gScore.BlueScore = 0
}
