package model

import (
	"math"
	"testing"

	"github.com/HackYourCareer/SmartKickers/internal/config"
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
		{
			name:               "should increment team white score by one",
			teamID:             config.TeamWhite,
			expectedBlueScore:  0,
			expectedWhiteScore: 1, expectedError: "",
		},
		{
			name:               "should increment team blue score by one",
			teamID:             config.TeamBlue,
			expectedBlueScore:  1,
			expectedWhiteScore: 0, expectedError: "",
		},
		{
			name:   "should cause an error when invalid team ID",
			teamID: -1, expectedBlueScore: 0,
			expectedWhiteScore: 0,
			expectedError:      "bad team ID",
		},
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
	game := &game{score: GameScore{1, 1}, scoreChannel: make(chan GameScore, 32)}

	type args struct {
		name               string
		teamID             int
		expectedBlueScore  int
		expectedWhiteScore int
		expectedError      string
	}

	tests := []args{
		{
			name:               "should decrement team white score by one",
			teamID:             config.TeamWhite,
			expectedBlueScore:  1,
			expectedWhiteScore: 0,
			expectedError:      "",
		},
		{
			name:               "should decrement team blue score by one",
			teamID:             config.TeamBlue,
			expectedBlueScore:  0,
			expectedWhiteScore: 1,
			expectedError:      "",
		},
		{
			name:               "should cause an error when invalid team ID",
			teamID:             -1,
			expectedBlueScore:  1,
			expectedWhiteScore: 1,
			expectedError:      "bad team ID",
		},
		{
			name:               "shoud not decrement the score",
			teamID:             config.TeamBlue,
			expectedBlueScore:  0,
			expectedWhiteScore: 0,
			expectedError:      "",
		},
	}

	for id, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.score.WhiteScore = 1
			game.score.BlueScore = 1

			if id == 3 {
				game.score.WhiteScore = 0
				game.score.BlueScore = 0
			}

			err := game.SubGoal(tt.teamID)
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

func TestUpdateShotsData(t *testing.T) {
	game := &game{}

	type args struct {
		name               string
		shot               Shot
		expectedCountWhite int
		expectedCountBlue  int
		expectedError      string
	}

	tests := []args{
		{
			name: "should increment team white shot count by one",
			shot: Shot{
				Speed: 15,
				Team:  1,
			},
			expectedCountWhite: 1,
			expectedCountBlue:  0,
			expectedError:      "",
		},
		{
			name: "should increment team blue shot count by one",
			shot: Shot{
				Speed: 15,
				Team:  2,
			},
			expectedCountWhite: 0,
			expectedCountBlue:  1,
			expectedError:      "",
		},
		{
			name: "should cause an error when invalid team ID",
			shot: Shot{
				Speed: 15,
				Team:  3,
			},
			expectedCountWhite: 0,
			expectedCountBlue:  0,
			expectedError:      "incorrect team ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.gameData.WhiteShotsCount = 0
			game.gameData.BlueShotsCount = 0

			err := game.UpdateShotsData(tt.shot)

			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

			assert.Equal(t, tt.expectedCountWhite, game.gameData.WhiteShotsCount)
			assert.Equal(t, tt.expectedCountBlue, game.gameData.BlueShotsCount)
		})
	}
}

func TestSaveFastestGoal(t *testing.T) {
	game := &game{
		gameData: GameStats{
			FastestShot: Shot{
				Speed: 18.98,
				Team:  1,
			},
		},
	}

	type args struct {
		name            string
		shotMsgIn       Shot
		expectedFastest Shot
	}

	tests := []args{
		{
			name: "Should save new fastest of team 1",
			shotMsgIn: Shot{
				Speed: 21.45,
				Team:  1,
			},
			expectedFastest: Shot{
				Speed: 21.45,
				Team:  1,
			},
		},
		{
			name: "Should save new fastest of team 2",
			shotMsgIn: Shot{
				Speed: 55.5555,
				Team:  2,
			},
			expectedFastest: Shot{
				Speed: 55.5555,
				Team:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.saveFastestShot(tt.shotMsgIn)

			assert.Equal(t, tt.expectedFastest, game.gameData.FastestShot)
		})
	}
}

func TestIncrementHeatmap(t *testing.T) {
	game := &game{}

	type args struct {
		name                 string
		xCord                float64
		yCord                float64
		startingHeatmapValue int
		expectedHeatmapValue int
		expectedError        string
	}
	tests := []args{
		{
			name:                 "should increment heatmap value on given cords by one",
			xCord:                0.1000,
			yCord:                0.9940,
			startingHeatmapValue: 0,
			expectedHeatmapValue: 1,
			expectedError:        "",
		},
		{
			name:                 "should increment heatmap value on given cords by one",
			xCord:                0.53126,
			yCord:                0.85485,
			startingHeatmapValue: 5,
			expectedHeatmapValue: 6,
			expectedError:        "",
		},
		{
			name:                 "should cause an error when index out of x range",
			xCord:                0.99500,
			yCord:                0.86236,
			startingHeatmapValue: 0,
			expectedHeatmapValue: 0,
			expectedError:        "x ball position index out of range",
		},
		{
			name:                 "should cause an error when index out of x range",
			xCord:                1.12634,
			yCord:                0.86236,
			startingHeatmapValue: 0,
			expectedHeatmapValue: 0,
			expectedError:        "x ball position index out of range",
		},
		{
			name:                 "should cause an error when index out of y range",
			xCord:                0.12634,
			yCord:                1.52563,
			startingHeatmapValue: 1,
			expectedHeatmapValue: 1,
			expectedError:        "y ball position index out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := int(math.Round(config.HeatmapAccuracy * tt.xCord))
			y := int(math.Round(config.HeatmapAccuracy * tt.yCord))

			err := game.IncrementHeatmap(tt.xCord, tt.yCord)
			if err == nil {
				game.gameData.Heatmap[x][y] = tt.startingHeatmapValue
				_ = game.IncrementHeatmap(tt.xCord, tt.yCord)
				assert.Equal(t, tt.expectedHeatmapValue, game.gameData.Heatmap[x][y])
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

		})
	}
}
