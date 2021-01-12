package simple

import "testing"

func Test_bisect(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name            string
		args            args
		wantTargetIndex int
	}{
		// TODO: Add test cases.
		{"test1", args{arr: []int{1, 2, 3, 4}, target: 2}, 1},
		{"test1", args{arr: []int{1, 3, 5, 7}, target: 5}, 2},
		{"test1", args{arr: []int{1, 3, 5, 7, 9}, target: 9}, 4},
		{"test1", args{arr: []int{1}, target: 1}, 0},
		{"test1", args{arr: []int{-10, -9, 10}, target: -9}, 1},
		{"test1", args{arr: []int{}, target: 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTargetIndex := bisect(tt.args.arr, tt.args.target); gotTargetIndex != tt.wantTargetIndex {
				t.Errorf("bisect() = %v, want %v", gotTargetIndex, tt.wantTargetIndex)
			}
		})
	}
}
