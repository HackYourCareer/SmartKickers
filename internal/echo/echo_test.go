package echo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_handleGoal(t *testing.T) {
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
			handleGoal(tt.args.team)
			assert.Equal(t, gameScore.BlueScore, tt.wantedBlueScore, "blue team score changes incorrectly")
			assert.Equal(t, gameScore.WhiteScore, tt.wantedWhiteScore, "white team score changes incorrectly")
		})
	}
}
