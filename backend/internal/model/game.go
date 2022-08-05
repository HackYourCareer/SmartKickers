package model

import "errors"

const (
	TeamWhite = 1
	TeamBlue  = 2
)

type Game interface {
	AddGoal(int) error
	ResetScore()
}

type game struct {
	score gameScore
}

type gameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

func NewGame() Game {
	return &game{}
}

func (g *game) ResetScore() {
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
}

func (g *game) AddGoal(teamID int) error {
	switch teamID {
	case TeamWhite:
		g.score.WhiteScore++
	case TeamBlue:
		g.score.BlueScore++
	default:
		return errors.New("bad team ID")
	}
	return nil
}
