package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResetScore(t *testing.T) {
	gameScore := &gameScore{3, 1}
	type args struct {
		score *gameScore
	}
	tests := []struct {
		name string
		args args
	}{
		{"Check if values are set to 0", args{gameScore}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameScore.ResetScore()
		})
		if tt.args.score.WhiteScore != 0 || tt.args.score.BlueScore != 0 {
			t.Errorf("Score did not reset. Goals white: %v, Goals blue: %v", tt.args.score.WhiteScore, tt.args.score.BlueScore)
		}
	}
}

func TestAddGoal(t *testing.T) {
	gameScore := &gameScore{3, 1}
	type args struct {
		team int
	}
	tests := []struct {
		name             string
		args             args
		wantedBlueScore  int
		wantedWhiteScore int
	}{
		{"should increment team white score by one", args{1}, 0, 1},
		{"should increment team blue score by one", args{2}, 1, 0},
		{"should not cause an error", args{-1}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameScore.WhiteScore = 0
			gameScore.BlueScore = 0
			gameScore.AddGoal(tt.args.team)
			assert.Equal(t, gameScore.BlueScore, tt.wantedBlueScore, "blue team score changes incorrectly")
			assert.Equal(t, gameScore.WhiteScore, tt.wantedWhiteScore, "white team score changes incorrectly")
		})
	}
}
