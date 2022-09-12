package adapter

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/HackYourCareer/SmartKickers/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpackShotMsg(t *testing.T) {
	type args struct {
		name           string
		msgIn          tableShotParams
		expectedMsgOut model.Shot
		expectedError  string
	}

	tests := []args{
		{
			name: "wrong area message, should return empty model.Shot and error",
			msgIn: tableShotParams{
				Speed:     25.5,
				StartArea: 1,
				EndArea:   17,
			},
			expectedMsgOut: model.Shot{},
			expectedError:  "couldn't decode teamID",
		},
		{
			name: "should return speed 25.5, team white and true",
			msgIn: tableShotParams{
				Speed:     25.5,
				StartArea: 20,
				EndArea:   27,
			},
			expectedMsgOut: model.Shot{
				Speed:      25.5,
				Team:       config.TeamWhite,
				ShotAtGoal: true,
			},
			expectedError: "",
		},
		{
			name: "should return speed 17.226, team blue and false",
			msgIn: tableShotParams{
				Speed:     17.226,
				StartArea: 24,
				EndArea:   27,
			},
			expectedMsgOut: model.Shot{
				Speed:      17.226,
				Team:       config.TeamBlue,
				ShotAtGoal: false,
			},
			expectedError: "",
		},
		{
			name:           "no message, should return empty shot message and error",
			msgIn:          tableShotParams{},
			expectedMsgOut: model.Shot{},
			expectedError:  "missing shot parameters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rawSlice []json.RawMessage
			jsn, err := json.Marshal(tt.msgIn)

			require.Nil(t, err)
			rawMsg := json.RawMessage(jsn)
			if tt.msgIn != (tableShotParams{}) {
				rawSlice = append(rawSlice, rawMsg)
			}

			tableShot := tableShotMsg{
				Params: rawSlice,
			}
			shotJSON, err := json.Marshal(tableShot)

			require.Nil(t, err)
			reader := bytes.NewReader(shotJSON)

			msg, err := UnpackShotMsg(reader)

			if tt.expectedError == "" {
				require.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
			assert.Equal(t, tt.expectedMsgOut, msg)
		})
	}
}

func TestDecodeTeam(t *testing.T) {
	type args struct {
		name          string
		areas         []int
		expectedTeam  int
		expectedError string
	}

	tests := []args{
		{
			name:          "should return team white",
			areas:         config.WhiteTeamArea[:],
			expectedTeam:  config.TeamWhite,
			expectedError: "",
		},
		{
			name:          "should return team blue",
			areas:         config.BlueTeamArea[:],
			expectedTeam:  config.TeamBlue,
			expectedError: "",
		},
		{
			name:          "should return error",
			areas:         []int{0, 1, 5, -15, 55, 700},
			expectedTeam:  0,
			expectedError: "couldn't decode teamID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, area := range tt.areas {
				team, err := decodeTeam(area)
				if tt.expectedError == "" {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, tt.expectedError)
				}
				assert.Equal(t, tt.expectedTeam, team)
			}
		})
	}
}

func Test_checkIfShotAtGoal(t *testing.T) {
	type args struct {
		areaID int
		teamID int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult bool
	}{
		{
			name: "should return true when blue team shots at white team goal",
			args: args{
				areaID: config.WhiteTeamGoalArea,
				teamID: config.TeamBlue,
			},
			expectedResult: true,
		},
		{
			name: "should return true when white team shots at blue team goal",
			args: args{
				areaID: config.BlueTeamGoalArea,
				teamID: config.TeamWhite,
			},
			expectedResult: true,
		},
		{
			name: "should return false when white team shots at their own goal",
			args: args{
				areaID: config.WhiteTeamGoalArea,
				teamID: config.TeamWhite,
			},
			expectedResult: false,
		},
		{
			name: "should return false when blue shot is blocked in the middle zone",
			args: args{
				areaID: config.WhiteTeamArea[1],
				teamID: config.TeamBlue,
			},
			expectedResult: false,
		},
		{
			name: "should return false when white shot is blocked in the middle zone",
			args: args{
				areaID: config.BlueTeamArea[2],
				teamID: config.TeamWhite,
			},
			expectedResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedResult, checkIfShotAtGoal(tt.args.areaID, tt.args.teamID))
		})
	}
}
