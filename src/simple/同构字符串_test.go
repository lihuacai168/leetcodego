package simple

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{s: "egg", t: "add"}, true},
		{"test2", args{s: "foo", t: "bar"}, false},
		{"test2", args{s: "bar", t: "foo"}, false},
		{"test3", args{s: "paper", t: "title"}, true},
		{"test4", args{s: "", t: ""}, true},
		{"test5", args{s: " ", t: " "}, true},
		{"test6", args{s: "ab", t: "aa"}, false},
		{"test7", args{s: "aa", t: "ab"}, false},
		{"test8", args{s: "aba", t: "baa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
