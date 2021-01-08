package medium

import (
	"reflect"
	"testing"
)

func Test_rotate1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"test1", args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 2}, []int{6, 7, 1, 2, 3, 4, 5}},
		{"test1", args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 4}, []int{4, 5, 6, 7, 1, 2, 3}},
		{"test1", args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 11}, []int{4, 5, 6, 7, 1, 2, 3}},
		{"test1", args{nums: []int{1, 2, 3}, k: 2}, []int{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate1(tt.args.nums, tt.args.k)
			if got := tt.args.nums; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
