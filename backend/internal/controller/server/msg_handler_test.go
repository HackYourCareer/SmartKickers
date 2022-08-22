package server

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/HackYourCareer/SmartKickers/internal/model"
	"github.com/stretchr/testify/assert"
)

type dispatcherMsg struct {
	MsgType string  `json:"type,omitempty"`
	TableID string  `json:"id,omitempty"`
	X       float64 `json:"x,omitempty"`
	Y       float64 `json:"y,omitempty"`
	Goal    int     `json:"goal,omitempty"`
}

type initialResponse struct {
	GameID string `json:"start,omitempty"`
}

func TestCreateResponse(t *testing.T) {
	type args struct {
		name          string
		msgIn         dispatcherMsg
		returnsNil    bool
		expectedError string
	}
	tests := []args{
		{
			name: "initial message",
			msgIn: dispatcherMsg{
				MsgType: "INITIAL",
				Goal:    0,
				TableID: "10",
				X:       0,
				Y:       0,
			},
			returnsNil:    false,
			expectedError: "",
		},
		{
			name: "goal message",
			msgIn: dispatcherMsg{
				MsgType: "",
				Goal:    2,
				TableID: "1",
				X:       0,
				Y:       0,
			},
			returnsNil:    true,
			expectedError: "",
		},
		{
			name: "position message",
			msgIn: dispatcherMsg{
				Goal:    0,
				TableID: "5",
				X:       14.3,
				Y:       1.25,
			},
			returnsNil:    true,
			expectedError: "unrecognized message type 0",
		},
		{
			name: "unexpected message",
			msgIn: dispatcherMsg{
				MsgType: "Something",
				Goal:    0,
				TableID: "Random",
				X:       14.3,
				Y:       1.25,
			},
			returnsNil:    true,
			expectedError: "unrecognized message type 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsn, _ := json.Marshal(tt.msgIn)
			reader := bytes.NewReader(jsn)

			s := New("0.0.0.0:3000", model.NewGame())

			response, err := s.createResponse(reader)

			expectedResponse, _ := json.Marshal(initialResponse{GameID: tt.msgIn.TableID})
			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
			if tt.returnsNil {
				assert.Nil(t, response)
			} else {
				assert.Equal(t, response, expectedResponse)
			}

		})
	}
}
