package simple

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{bills: []int{5, 5, 5, 10, 20}}, true},
		{"test2", args{bills: []int{5, 5, 10}}, true},
		{"test3", args{bills: []int{10, 10}}, false},
		{"test4", args{bills: []int{5, 5, 10, 10, 20}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
