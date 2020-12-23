package simple

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{s: "leetcode"}, 0},
		{"test2", args{s: "loveleetcode"}, 2},
		{"test3", args{s: "abcabc"}, -1},
		{"test4", args{s: "aaa"}, -1},
		{"test5", args{s: "aba"}, 1},
		{"test5", args{s: "dddccdbba"}, 8},
		{"test5", args{s: ""}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
