package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResetScore(t *testing.T) {
	game := &game{score: GameScore{3, 1}, scoreChannel: make(chan GameScore, 32)}

	game.ResetScore()
	resultScore := <-game.scoreChannel

	if resultScore.BlueScore != 0 || resultScore.WhiteScore != 0 {
		t.Errorf("Score did not reset. Goals white: %v, Goals blue: %v", game.score.WhiteScore, game.score.BlueScore)
	}
}

func TestAddGoal(t *testing.T) {
	game := &game{score: GameScore{3, 1}, scoreChannel: make(chan GameScore, 32)}

	type args struct {
		name               string
		teamID             int
		expectedBlueScore  int
		expectedWhiteScore int
		expectedError      string
	}
	tests := []args{
		{name: "should increment team white score by one", teamID: TeamWhite, expectedBlueScore: 0, expectedWhiteScore: 1, expectedError: ""},
		{name: "should increment team blue score by one", teamID: TeamBlue, expectedBlueScore: 1, expectedWhiteScore: 0, expectedError: ""},
		{name: "should cause an error when invalid team ID", teamID: -1, expectedBlueScore: 0, expectedWhiteScore: 0, expectedError: "bad team ID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			game.score.WhiteScore = 0
			game.score.BlueScore = 0
			err := game.AddGoal(tt.teamID)
			if err == nil {
				resultScore := <-game.scoreChannel

				assert.Equal(t, resultScore.BlueScore, tt.expectedBlueScore, "blue team score changes incorrectly")
				assert.Equal(t, resultScore.WhiteScore, tt.expectedWhiteScore, "white team score changes incorrectly")
			}

			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}

func TestGameSubGoal(t *testing.T) {
	game := &game{score: GameScore{0, 2}}

	type args struct {
		name               string
		teamID             int
		expectedBlueScore  int
		expectedWhiteScore int
		expectedError      string
	}
	tests := []args{
		{name: "should decrement team white score by one", teamID: TeamWhite, expectedBlueScore: 2, expectedWhiteScore: 0, expectedError: ""},
		{name: "should decrement team blue score by one", teamID: TeamBlue, expectedBlueScore: 1, expectedWhiteScore: 1, expectedError: ""},
		{name: "should not decrement team blue score by one", teamID: TeamWhite, expectedBlueScore: 2, expectedWhiteScore: 0, expectedError: ""},
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
