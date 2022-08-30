package model

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/HackYourCareer/SmartKickers/internal/config"

	log "github.com/sirupsen/logrus"
)

type Game interface {
	AddGoal(int) error
	ResetScore()
	GetScore() GameScore
	GetScoreChannel() chan GameScore
	SubGoal(int) error
	UpdateManualGoals(int, string)
	UpdateShotsData(Shot) error
	GetShotsData() ShotsData
	IncrementHeatmap(float64, float64) error
}

type game struct {
	score        GameScore
	shotsData    ShotsData
	scoreChannel chan GameScore
	manualGoals  map[int]map[string]int
	m            sync.RWMutex
	heatmap      [config.HeatmapAccuracy][config.HeatmapAccuracy]int
}

type GameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

type ShotsData struct {
	WhiteCount       int
	BlueCount        int
	Fastest          Shot
	BlueAtGoalCount  int
	WhiteAtGoalCount int
}

type Shot struct {
	Speed      float64
	Team       int
	ShotAtGoal bool
}

func NewGame() Game {
	return &game{
		scoreChannel: make(chan GameScore, 32),
		manualGoals: map[int]map[string]int{
			config.TeamWhite: {
				config.ActionAdd:      0,
				config.ActionSubtract: 0,
			},
			config.TeamBlue: {
				config.ActionAdd:      0,
				config.ActionSubtract: 0,
			},
		},
	}
}

func (g *game) ResetScore() {
	log.Trace("mutex lock: ResetScore")
	g.m.Lock()
	defer g.m.Unlock()
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
	g.scoreChannel <- g.score
}

func (g *game) AddGoal(teamID int) error {
	log.Trace("mutex lock: AddGoal")
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
	log.Trace("mutex lock: GetScore")
	g.m.RLock()
	defer g.m.RUnlock()

	return g.score
}

func (g *game) GetScoreChannel() chan GameScore {
	return g.scoreChannel
}

func (g *game) SubGoal(teamID int) error {
	log.Trace("mutex lock: SubGoal")
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

func (g *game) UpdateShotsData(shot Shot) error {
	log.Trace("mutex lock: UpdateRecordedShots")
	g.m.Lock()
	defer g.m.Unlock()

	switch shot.Team {
	case config.TeamWhite:
		g.shotsData.WhiteCount++
		if shot.ShotAtGoal {
			g.shotsData.WhiteAtGoalCount++
		}
	case config.TeamBlue:
		g.shotsData.BlueCount++
		if shot.ShotAtGoal {
			g.shotsData.BlueAtGoalCount++
		}
	default:
		return fmt.Errorf("incorrect team ID")
	}

	if g.isFastestShot(shot.Speed) {
		g.saveFastestShot(shot)
	}

	return nil
}

func (g *game) isFastestShot(speed float64) bool {
	return g.shotsData.Fastest.Speed < speed
}

func (g *game) saveFastestShot(shot Shot) {
	g.shotsData.Fastest.Speed = shot.Speed
	g.shotsData.Fastest.Team = shot.Team
}

func (g *game) GetShotsData() ShotsData {
	log.Trace("mutex lock: GetRecordedShots")
	g.m.RLock()
	defer g.m.RUnlock()

	return g.shotsData
}

func (g *game) UpdateManualGoals(teamID int, action string) {
	g.m.Lock()
	defer g.m.Unlock()
	g.manualGoals[teamID][action]++
}

func (g *game) IncrementHeatmap(xCord float64, yCord float64) error {
	log.Trace("mutex lock: WriteToHeatmap")
	g.m.Lock()
	defer g.m.Unlock()

	x := int(math.Round(config.HeatmapAccuracy * xCord))
	y := int(math.Round(config.HeatmapAccuracy * yCord))
	heatmapUpperBound := config.HeatmapAccuracy - 1
	if x > heatmapUpperBound || x < 0 {
		return errors.New("x ball position index out of range")
	}
	if y > heatmapUpperBound || y < 0 {
		return errors.New("y ball position index out of range")
	}
	g.heatmap[x][y]++
	return nil
}
