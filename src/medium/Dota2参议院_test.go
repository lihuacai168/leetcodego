package medium

import "testing"

func Test_predictPartyVictory(t *testing.T) {
	type args struct {
		senate string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{senate: "RD"}, "Radiant"},
		{"test2", args{senate: "RDD"}, "Dire"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictPartyVictory(tt.args.senate); got != tt.want {
				t.Errorf("predictPartyVictory() = %v, want %v", got, tt.want)
			}
		})
	}
}
