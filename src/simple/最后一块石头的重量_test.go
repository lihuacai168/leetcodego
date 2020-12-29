package simple

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{stones: []int{2, 7, 4, 1, 8, 1}}, 1},
		{"test2", args{stones: []int{2}}, 2},
		{"test3", args{stones: []int{2, 2}}, 0},
		{"test4", args{stones: []int{2, 2, 1}}, 1},
		{"test5", args{stones: []int{1, 3}}, 2},
		{"test6", args{stones: []int{2, 3, 7}}, 2},
		{"test7", args{stones: []int{3, 7, 8}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
