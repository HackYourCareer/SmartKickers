package adapter

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessageCategory(t *testing.T) {

	type args struct {
		name             string
		msg              dispatcherMsg
		expectedCategory MsgCategory
	}
	tests := []args{
		{
			name: "initial message, should return MsgInitial",
			msg: dispatcherMsg{
				MsgType: "INITIAL",
				Goal:    0,
				X:       0,
				Y:       0,
			},
			expectedCategory: MsgInitial},
		{
			name: "goal message, should return MsgGoal",
			msg: dispatcherMsg{
				MsgType: "",
				Goal:    1,
				X:       0,
				Y:       0,
			},
			expectedCategory: MsgGoal},
		{
			name: "position message, should return MsgNone",
			msg: dispatcherMsg{
				MsgType: "",
				Goal:    0,
				X:       4,
				Y:       2.22,
			},
			expectedCategory: MsgNone},
		{
			name: "unexpected message, should return MsgNone",
			msg: dispatcherMsg{
				MsgType: "SOMETHING",
				Goal:    0,
				X:       15,
				Y:       0},
			expectedCategory: MsgNone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category := tt.msg.getMessageCategory()
			assert.Equal(t, category, tt.expectedCategory)
		})
	}
}

func TestUnpack(t *testing.T) {
	type args struct {
		name           string
		msgIn          dispatcherMsg
		ExpectedMsgOut Message
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
			ExpectedMsgOut: Message{
				Category: MsgInitial,
				TableID:  "10",
				Team:     0,
			},
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
			ExpectedMsgOut: Message{
				Category: MsgGoal,
				TableID:  "1",
				Team:     2,
			},
		},
		{
			name: "position message",
			msgIn: dispatcherMsg{
				MsgType: "",
				Goal:    0,
				TableID: "5",
				X:       14.3,
				Y:       1.25,
			},
			ExpectedMsgOut: Message{
				Category: MsgNone,
				TableID:  "5",
				Team:     0,
			},
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
			ExpectedMsgOut: Message{
				Category: MsgNone,
				TableID:  "Random",
				Team:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsn, _ := json.Marshal(tt.msgIn)
			reader := bytes.NewReader(jsn)

			msg, _ := UnpackDispatcherMsg(reader)
			assert.Equal(t, msg, tt.ExpectedMsgOut)
		})
	}
}
