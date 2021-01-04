package simple

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{n: 0}, 0},
		{"test1", args{n: 1}, 1},
		{"test2", args{n: 2}, 1},
		{"test3", args{n: 3}, 2},
		{"test4", args{n: 4}, 3},
		{"test5", args{n: 5}, 5},
		{"test6", args{n: 6}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
