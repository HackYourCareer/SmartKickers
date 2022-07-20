package messages

import (
	"testing"
)

func Test_resetScore(t *testing.T) {
	var gameScore *GameScore = &GameScore{2, 0}
	type args struct {
		score *GameScore
	}
	tests := []struct {
		name string
		args args
	}{
		{"Check if values are set to 0", args{gameScore}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameScore.ResetScore()
		})
		if *&tt.args.score.WhiteScore != 0 || *&tt.args.score.BlueScore != 0 {
			t.Errorf("Score did not reset. Goals white: %v, Goals blue: %v", *&tt.args.score.WhiteScore, *&tt.args.score.BlueScore)
		}

	}
}
