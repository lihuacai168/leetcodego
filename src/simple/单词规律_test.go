package simple

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"abba", args{"abba", "dog cat cat dog"}, true},
		{"abbaFalse", args{"abba", "dog cat cat dog1"}, false},
		{"aaaa", args{"aaaa", "dog cat cat dog1"}, false},
		{"abbaDog", args{"abba", "cat cat cat cat"}, false},
		{"aaa", args{"aaa", "cat cat cat cat"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
