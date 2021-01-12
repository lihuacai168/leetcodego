package simple

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test1", args{nums: []int{0, 1, 2, 4, 5, 7}}, []string{"0->2", "4->5", "7"}},
		{"test2", args{nums: []int{0, 2, 3, 4, 6, 8, 9}}, []string{"0", "2->4", "6", "8->9"}},
		{"test3", args{nums: []int{-1}}, []string{"-1"}},
		{"test4", args{nums: []int{0}}, []string{"0"}},
		{"test5", args{nums: []int{}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
