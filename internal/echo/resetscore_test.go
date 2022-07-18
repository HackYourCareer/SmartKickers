package echo

import "testing"

func Test_resetScore(t *testing.T) {
	gWhite := 4
	gBlue := 20
	type args struct {
		gW *int
		gB *int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{&gWhite, &gBlue}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetScore(tt.args.gW, tt.args.gB)
		})
		if *tt.args.gW != 0 || *tt.args.gB != 0 {
			t.Errorf("Score did not reset. Goals white: %v, Goals blue: %v", *tt.args.gW, *tt.args.gB)
		}

	}
}
