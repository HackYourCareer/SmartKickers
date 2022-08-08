package model

import (
	"errors"
	"sync"
)

const (
	teamWhite = 1
	teamBlue  = 2
)

type Game interface {
	AddGoal(int) error
	ResetScore()
	GetScore() gameScore
	GetScoreChannel() chan gameScore
	SubGoal(int) error
}

type game struct {
	score        gameScore
	scoreChannel chan gameScore
	m            sync.RWMutex
}

type gameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

func NewGame() Game {
	return &game{scoreChannel: make(chan gameScore)}
}

func (g *game) ResetScore() {
	g.m.Lock()
	defer g.m.Unlock()
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
	g.scoreChannel <- g.score
}

func (g *game) AddGoal(teamID int) error {
	g.m.Lock()
	defer g.m.Unlock()
	switch teamID {
	case teamWhite:
		g.score.WhiteScore++
	case teamBlue:
		g.score.BlueScore++
	default:
		return errors.New("bad team ID")
	}
	g.scoreChannel <- g.score
	return nil
}

func (g *game) GetScore() gameScore {
	return g.score
}

func (g *game) GetScoreChannel() chan gameScore {
	return g.scoreChannel
}

func (g *game) SubGoal(teamID int) error {
	g.m.Lock()
	defer g.m.Unlock()
	switch teamID {
	case teamWhite:
		if g.score.WhiteScore > 0 {
			g.score.WhiteScore--
		}
	case teamBlue:
		if g.score.BlueScore > 0 {
			g.score.BlueScore--
		}
	default:
		return errors.New("bad team ID")
	}
	g.scoreChannel <- g.score
	return nil

}
