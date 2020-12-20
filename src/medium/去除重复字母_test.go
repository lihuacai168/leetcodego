package medium

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{"bcabc"}, "abc"},
		{"test2", args{"cbacdcbc"}, "acdb"},
		{"test2", args{"cdadabcc"}, "adbc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
