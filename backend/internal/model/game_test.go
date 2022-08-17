package model

import (
	"testing"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
	"github.com/stretchr/testify/assert"
)

func TestResetScore(t *testing.T) {
	game := &game{score: GameScore{3, 1}}

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
	game := &game{score: GameScore{3, 1}}

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
		{
			name:               "should increment team white score by one",
			teamID:             config.TeamWhite,
			expectedBlueScore:  0,
			expectedWhiteScore: 1,
			expectedError:      "",
		},
		{
			name:               "should increment team blue score by one",
			teamID:             config.TeamBlue,
			expectedBlueScore:  1,
			expectedWhiteScore: 0,
			expectedError:      "",
		},
		{
			name:               "should cause an error when invalid team ID",
			teamID:             -1,
			expectedBlueScore:  0,
			expectedWhiteScore: 0,
			expectedError:      "bad team ID",
		},
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

func TestUpdateShotsData(t *testing.T) {
	game := &game{}

	type args struct {
		name               string
		shot               adapter.ShotMessage
		expectedCountWhite int
		expectedCountBlue  int
		expectedError      string
	}
	tests := []args{
		{
			name: "should increment team white shot count by one",
			shot: adapter.ShotMessage{
				Speed: 15,
				Team:  1,
			},
			expectedCountWhite: 1,
			expectedCountBlue:  0,
			expectedError:      "",
		},
		{
			name: "should increment team blue shot count by one",
			shot: adapter.ShotMessage{
				Speed: 15,
				Team:  2,
			},
			expectedCountWhite: 0,
			expectedCountBlue:  1,
			expectedError:      "",
		},
		{
			name: "should cause an error when invalid team ID",
			shot: adapter.ShotMessage{
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
			game.shotsData.WhiteCount = 0
			game.shotsData.BlueCount = 0

			err := game.UpdateShotsData(tt.shot)

			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}

			assert.Equal(t, tt.expectedCountWhite, game.shotsData.WhiteCount)
			assert.Equal(t, tt.expectedCountBlue, game.shotsData.BlueCount)
		})
	}
}

func TestSaveFastestGoal(t *testing.T) {

	game := &game{
		shotsData: ShotsData{
			Fastest: adapter.ShotMessage{
				Speed: 18.98,
				Team:  1,
			},
		},
	}

	type args struct {
		name            string
		shotMsgIn       adapter.ShotMessage
		expectedFastest adapter.ShotMessage
	}

	tests := []args{
		{
			name: "Should save new fastest of team 1",
			shotMsgIn: adapter.ShotMessage{
				Speed: 21.45,
				Team:  1,
			},
			expectedFastest: adapter.ShotMessage{
				Speed: 21.45,
				Team:  1,
			},
		},
		{
			name: "Should save new fastest of team 2",
			shotMsgIn: adapter.ShotMessage{
				Speed: 55.5555,
				Team:  2,
			},
			expectedFastest: adapter.ShotMessage{
				Speed: 55.5555,
				Team:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.saveFastestShot(tt.shotMsgIn)

			assert.Equal(t, tt.expectedFastest, game.shotsData.Fastest)
		})
	}

}
