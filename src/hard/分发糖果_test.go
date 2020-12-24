package hard

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{ratings: []int{1, 0, 2}}, 5},
		{"test1", args{ratings: []int{1, 3, 2, 2, 1}}, 7},
		{"test1", args{ratings: []int{1, 2, 2}}, 4},
		{"test1", args{ratings: []int{1, 1, 1}}, 6},
		{"test1", args{ratings: []int{1, 1}}, 3},
		{"test1", args{ratings: []int{}}, 0},
		{"test1", args{ratings: []int{111}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
