package simple

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{s: "abbxxxxzzy"}, [][]int{{3, 6}}},
		{"test1", args{s: "abc"}, [][]int{}},
		{"test1", args{s: "aaa"}, [][]int{{0, 2}}},
		{"test1", args{s: "aaaa"}, [][]int{{0, 3}}},
		{"test1", args{s: "aba"}, [][]int{}},
		{"test1", args{s: "bababbaaab"}, [][]int{{6, 8}}},
		{"test1", args{s: "abcdddeeeeaabbbcd"}, [][]int{{3, 5}, {6, 9}, {12, 14}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
