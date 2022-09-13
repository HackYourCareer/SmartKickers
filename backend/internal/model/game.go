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
	ResetStats()
	GetScore() GameScore
	GetScoreChannel() chan GameScore
	SubGoal(int) error
	UpdateManualGoals(int, string)
	UpdateShotsData(Shot) error
	GetGameStats() GameStats
	GetHeatmap() [config.HeatmapAccuracy][config.HeatmapAccuracy]int
	IncrementHeatmap(float64, float64) error
}

type game struct {
	score        GameScore
	gameData     GameStats
	heatmap      [config.HeatmapAccuracy][config.HeatmapAccuracy]int
	scoreChannel chan GameScore
	m            sync.RWMutex
}

type GameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

type GameStats struct {
	WhiteShotsCount  int                    `json:"whiteShotsCount"`
	BlueShotsCount   int                    `json:"blueShotsCount"`
	FastestShot      Shot                   `json:"fastestShot"`
	ManualGoals      map[int]map[string]int `json:"manualGoals"`
	BlueAtGoalCount  int
	WhiteAtGoalCount int
}

type Shot struct {
	Speed      float64 `json:"speed"`
	Team       int     `json:"team"`
	ShotAtGoal bool    `json:"shotAtGoal"`
}

func NewGame() Game {
	return &game{
		scoreChannel: make(chan GameScore, 32),
		gameData: GameStats{
			ManualGoals: map[int]map[string]int{
				config.TeamWhite: {
					config.ActionAdd:      0,
					config.ActionSubtract: 0,
				},
				config.TeamBlue: {
					config.ActionAdd:      0,
					config.ActionSubtract: 0,
				},
			},
		},
	}
}

func (g *game) ResetStats() {
	log.Trace("mutex lock: ResetStats")
	g.m.Lock()
	defer g.m.Unlock()
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
	g.scoreChannel <- g.score
	g.gameData = GameStats{
		ManualGoals: map[int]map[string]int{
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
	g.heatmap = [config.HeatmapAccuracy][config.HeatmapAccuracy]int{}
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
	log.Trace("mutex lock: UpdateShotsData")
	g.m.Lock()
	defer g.m.Unlock()

	switch shot.Team {
	case config.TeamWhite:
		g.gameData.WhiteShotsCount++
		if shot.ShotAtGoal {
			g.gameData.WhiteAtGoalCount++
		}
	case config.TeamBlue:
		g.gameData.BlueShotsCount++
		if shot.ShotAtGoal {
			g.gameData.BlueAtGoalCount++
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
	return g.gameData.FastestShot.Speed < speed
}

func (g *game) saveFastestShot(shot Shot) {
	g.gameData.FastestShot.Speed = shot.Speed
	g.gameData.FastestShot.Team = shot.Team
}

func (g *game) GetGameStats() GameStats {
	log.Trace("mutex lock: GetGameStats")
	g.m.RLock()
	defer g.m.RUnlock()

	return g.gameData
}

func (g *game) UpdateManualGoals(teamID int, action string) {
	g.m.Lock()
	defer g.m.Unlock()
	g.gameData.ManualGoals[teamID][action]++
}

func (g *game) IncrementHeatmap(xCord float64, yCord float64) error {
	log.Trace("mutex lock: IncrementHeatmap")
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

func (g *game) GetHeatmap() [config.HeatmapAccuracy][config.HeatmapAccuracy]int {
	log.Trace("mutex lock: GetHeatmap")
	g.m.RLock()
	defer g.m.RUnlock()

	return g.heatmap
}
