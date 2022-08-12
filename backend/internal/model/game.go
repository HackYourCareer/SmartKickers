package model

import (
	"errors"
	"fmt"
	"sync"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
)

type Game interface {
	AddGoal(int) error
	ResetScore()
	GetScore() GameScore
	GetScoreChannel() chan GameScore
	SubGoal(int) error
	IsFastestShot(float64) bool
	UpdateRecordedShots(adapter.ShotMessage) error
	GetRecordedShots() Shots
}

type game struct {
	score         GameScore
	recordedShots Shots
	scoreChannel  chan GameScore
	m             sync.RWMutex
}

type GameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

type Shots struct {
	White   int
	Blue    int
	Fastest adapter.ShotMessage
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
	case config.TeamWhite:
		g.score.WhiteScore++
	case config.TeamBlue:
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
	case config.TeamWhite:
		if g.score.WhiteScore > 0 {
			g.score.WhiteScore--
		}
	case config.TeamBlue:
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

	return g.recordedShots.Fastest.Speed < speed
}

func (g *game) UpdateRecordedShots(shot adapter.ShotMessage) error {
	g.m.Lock()
	defer g.m.Unlock()

	switch shot.Team {
	case config.TeamWhite:
		g.recordedShots.White++
	case config.TeamBlue:
		g.recordedShots.Blue++
	default:
		return fmt.Errorf("incorrect team ID")
	}

	if g.IsFastestShot(shot.Speed) {
		g.saveFastestShot(shot)
	}

	return nil
}

func (g *game) saveFastestShot(shot adapter.ShotMessage) {
	g.m.Lock()
	defer g.m.Unlock()
	g.recordedShots.Fastest.Speed = shot.Speed
	g.recordedShots.Fastest.Team = shot.Team
}

func (g *game) GetRecordedShots() Shots {
	g.m.RLock()
	defer g.m.RUnlock()

	return g.recordedShots
}
