package simple

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{nums: []int{1, 2, 3}}, false},
		{"test2", args{nums: []int{1}}, false},
		{"test3", args{nums: []int{}}, false},
		{"test4", args{nums: []int{1, 2, 2, 3}}, true},
		{"test4", args{nums: []int{1, 2, 3, 4, 5, 6, 6}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
