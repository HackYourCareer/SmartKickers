package model

import (
	"errors"
	"sync"

	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
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
	IsFastestShot(float64) bool
	SaveFastestShot(adapter.ShotMessage)
}

type game struct {
	score        GameScore
	scoreChannel chan GameScore
	m            sync.RWMutex
	fastestShot  adapter.ShotMessage
}

type GameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
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

func (g *game) IsFastestShot(speed float64) bool {
	g.m.RLock()
	defer g.m.RUnlock()

	return g.fastestShot.Speed < speed
}

func (g *game) SaveFastestShot(msg adapter.ShotMessage) {
	g.m.Lock()
	defer g.m.Unlock()
	g.fastestShot.Speed = msg.Speed
	g.fastestShot.Team = msg.Team
}
