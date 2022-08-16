package adapter

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestUnpackShotMsg(t *testing.T) {
	type args struct {
		name           string
		msgIn          tableShotParams
		expectedMsgOut ShotMessage
		expectedError  string
	}
	tests := []args{
		{
			name: "wrong area message, should return empty ShotMessage and error",
			msgIn: tableShotParams{
				Speed:     25.5,
				StartArea: 1,
			},
			expectedMsgOut: ShotMessage{},
			expectedError:  "couldn't decode teamID",
		},
		{
			name: "should return speed 25.5 and team white",
			msgIn: tableShotParams{
				Speed:     25.5,
				StartArea: 20,
			},
			expectedMsgOut: ShotMessage{
				Speed: 25.5,
				Team:  config.TeamWhite,
			},
			expectedError: "",
		},
		{
			name: "should return speed 17.226 and team blue",
			msgIn: tableShotParams{
				Speed:     17.226,
				StartArea: 24,
			},
			expectedMsgOut: ShotMessage{
				Speed: 17.226,
				Team:  config.TeamBlue,
			},
			expectedError: "",
		},
		{
			name:           "no message, should return empty shot message and error",
			msgIn:          tableShotParams{},
			expectedMsgOut: ShotMessage{},
			expectedError:  "missing shot parameters",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rawSlice []json.RawMessage
			jsn, _ := json.Marshal(tt.msgIn)
			rawMsg := json.RawMessage(jsn)
			if tt.msgIn != (tableShotParams{}) {
				rawSlice = append(rawSlice, rawMsg)
			}

			tableShot := tableShotMsg{
				Params: rawSlice,
			}
			shotJSON, _ := json.Marshal(tableShot)
			reader := bytes.NewReader(shotJSON)

			msg, err := UnpackShotMsg(reader)

			if tt.expectedError == "" {
				assert.Nil(t, err)
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
