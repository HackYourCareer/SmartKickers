package model

import (
	"errors"
	"sync"
)

const (
	TeamWhite = 1
	TeamBlue  = 2
)

type Game interface {
	AddGoal(int) error
	ResetScore()
	GetScore() GameScore
	GetScoreChannel() chan GameScore
	SubGoal(int) error
	UpdateManualGoals(int, bool) error
}

type game struct {
	score        GameScore
	scoreChannel chan GameScore
	manualGoals  ManualGoals
	m            sync.RWMutex
}

type GameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

type ManualGoals struct {
	AddedBlue       int
	SubtractedBlue  int
	AddedWhite      int
	SubtractedWhite int
}

func NewGame() Game {
	return &game{scoreChannel: make(chan GameScore, 32)}
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
	case TeamWhite:
		g.score.WhiteScore++
	case TeamBlue:
		g.score.BlueScore++
	default:
		return errors.New("bad team ID")
	}
	g.scoreChannel <- g.score
	return nil
}

func (g *game) GetScore() GameScore {
	g.m.RLock()
	defer g.m.RUnlock()
	return g.score
}

func (g *game) GetScoreChannel() chan GameScore {
	return g.scoreChannel
}

func (g *game) SubGoal(teamID int) error {
	g.m.Lock()
	defer g.m.Unlock()
	switch teamID {
	case TeamWhite:
		if g.score.WhiteScore > 0 {
			g.score.WhiteScore--
		}
	case TeamBlue:
		if g.score.BlueScore > 0 {
			g.score.BlueScore--
		}
	default:
		return errors.New("bad team ID")
	}
	g.scoreChannel <- g.score
	return nil

}

func (g *game) UpdateManualGoals(teamID int, addGoal bool) error {
	g.m.Lock()
	defer g.m.Unlock()
	if addGoal {
		switch teamID {
		case TeamWhite:
			g.manualGoals.AddedWhite++
		case TeamBlue:
			g.manualGoals.AddedBlue++
		default:
			return errors.New("bad team ID")
		}
		return nil
	}
	switch teamID {
	case TeamWhite:
		g.manualGoals.SubtractedWhite++
	case TeamBlue:
		g.manualGoals.SubtractedBlue++
	default:
		return errors.New("bad team ID")
	}

	return nil
}
