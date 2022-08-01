package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"errors"
)

func TestResetScore(t *testing.T) {
	game := &Game{gameScore{3, 1}}

	type args struct {
		name               string
		expectedBlueScore  int
		expectedWhiteScore int
	}

	tests := []args{
		{name: "Check if values are set to 0", expectedBlueScore: 0, expectedWhiteScore: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.ResetScore()
		})
		if tt.expectedWhiteScore != 0 || tt.expectedBlueScore != 0 {
			t.Errorf("Score did not reset. Goals white: %v, Goals blue: %v", tt.expectedWhiteScore, tt.expectedBlueScore)
		}
	}
}

func TestAddGoal(t *testing.T) {
	game := &Game{gameScore{3, 1}}

	type args struct {
		name               string
		teamID             int
		expectedBlueScore  int
		expectedWhiteScore int
		expectedError      error
	}
	tests := []args{
		{name: "should increment team white score by one", teamID: 1, expectedBlueScore: 0, expectedWhiteScore: 1, expectedError: nil},
		{name: "should increment team blue score by one", teamID: 2, expectedBlueScore: 1, expectedWhiteScore: 0, expectedError: nil},
		{name: "should cause an error when invalid team ID", teamID: -1, expectedBlueScore: 0, expectedWhiteScore: 0, expectedError: errors.New("bad team ID")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.score.WhiteScore = 0
			game.score.BlueScore = 0
			game.AddGoal(tt.teamID)
			assert.Equal(t, game.score.BlueScore, tt.expectedBlueScore, "blue team score changes incorrectly")
			assert.Equal(t, game.score.WhiteScore, tt.expectedWhiteScore, "white team score changes incorrectly")
		})
	}
}
