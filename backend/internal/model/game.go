package model

import (
	"errors"
	"fmt"
	"sync"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"

	log "github.com/sirupsen/logrus"
)

type Game interface {
	AddGoal(int) error
	ResetScore()
	GetScore() GameScore
	GetScoreChannel() chan GameScore
	SubGoal(int) error
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
	log.Debug("mutex lock: ResetScore")
	g.m.Lock()
	defer g.m.Unlock()
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
	g.scoreChannel <- g.score
}

func (g *game) AddGoal(teamID int) error {
	log.Debug("mutex lock: AddGoal")
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
	log.Debug("mutex lock: GetScore")
	g.m.RLock()
	defer g.m.RUnlock()

	return g.score
}

func (g *game) GetScoreChannel() chan GameScore {
	return g.scoreChannel
}

func (g *game) SubGoal(teamID int) error {
	log.Debug("mutex lock: SubGoal")
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

func (g *game) UpdateRecordedShots(shot adapter.ShotMessage) error {
	log.Debug("mutex lock: UpdateRecordedShots")
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

	if g.isFastestShot(shot.Speed) {
		g.saveFastestShot(shot)
	}

	return nil
}

func (g *game) isFastestShot(speed float64) bool {
	return g.recordedShots.Fastest.Speed < speed
}

func (g *game) saveFastestShot(shot adapter.ShotMessage) {
	g.recordedShots.Fastest.Speed = shot.Speed
	g.recordedShots.Fastest.Team = shot.Team
}

func (g *game) GetRecordedShots() Shots {
	log.Debug("mutex lock: GetRecordedShots")
	g.m.RLock()
	defer g.m.RUnlock()

	return g.recordedShots
}
