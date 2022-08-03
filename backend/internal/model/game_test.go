package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResetScore(t *testing.T) {
	game := &game{gameScore{3, 1}}

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
	game := &game{gameScore{3, 1}}

	type args struct {
		name               string
		teamID             int
		expectedBlueScore  int
		expectedWhiteScore int
		expectedError      string
	}
	tests := []args{
		{name: "should increment team white score by one", teamID: teamWhite, expectedBlueScore: 0, expectedWhiteScore: 1, expectedError: ""},
		{name: "should increment team blue score by one", teamID: teamBlue, expectedBlueScore: 1, expectedWhiteScore: 0, expectedError: ""},
		{name: "should cause an error when invalid team ID", teamID: -1, expectedBlueScore: 0, expectedWhiteScore: 0, expectedError: "bad team ID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.score.WhiteScore = 0
			game.score.BlueScore = 0
			err := game.AddGoal(tt.teamID)
			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
			assert.Equal(t, game.score.BlueScore, tt.expectedBlueScore, "blue team score changes incorrectly")
			assert.Equal(t, game.score.WhiteScore, tt.expectedWhiteScore, "white team score changes incorrectly")
		})
	}
}

func Test_game_SubGoal(t *testing.T) {
	game := &game{gameScore{0, 2}}

	type args struct {
		name               string
		teamID             int
		expectedBlueScore  int
		expectedWhiteScore int
		expectedError      string
	}
	tests := []args{
		{name: "should decrement team white score by one", teamID: teamWhite, expectedBlueScore: 2, expectedWhiteScore: 0, expectedError: ""},
		{name: "should decrement team blue score by one", teamID: teamBlue, expectedBlueScore: 1, expectedWhiteScore: 1, expectedError: ""},
		{name: "should not decrement team blue score by one", teamID: teamWhite, expectedBlueScore: 2, expectedWhiteScore: 0, expectedError: ""},
		{name: "should cause an error when invalid team ID", teamID: -1, expectedBlueScore: 2, expectedWhiteScore: 1, expectedError: "bad team ID"},
	}

	for id, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.score.WhiteScore = 1
			game.score.BlueScore = 2
			if id == 2 {
				game.score.WhiteScore = 0
			}
			err := game.SubGoal(tt.teamID)
			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
			assert.Equal(t, game.score.BlueScore, tt.expectedBlueScore, "blue team score changes incorrectly")
			assert.Equal(t, game.score.WhiteScore, tt.expectedWhiteScore, "white team score changes incorrectly")
		})
	}
}
