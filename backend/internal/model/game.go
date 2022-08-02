package model

import "errors"

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

func NewGame() *game {
	return &game{}
}

func (g *game) ResetScore() {
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
}

func (g *game) AddGoal(teamID int) error {
	switch teamID {
	case 1:
		g.score.WhiteScore++
	case 2:
		g.score.BlueScore++
	default:
		return errors.New("bad team ID")
	}
	return nil
}
