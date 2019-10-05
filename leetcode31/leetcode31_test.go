package leetcode31

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[1,2,3], expect:[1,3,2]",
			args: args{
				nums: []int{1, 2, 3},
			},
		},
		{
			name: "[1,3,2], expect:[2,1,3]",
			args: args{
				nums: []int{1, 3, 2},
			},
		},
		{
			name: "[1,1,5], expect:[1,5,1]",
			args: args{
				nums: []int{1, 1, 5},
			},
		},
		{
			name: "[2,3,1], expect:[3,1,2]",
			args: args{
				nums: []int{2, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
		})
	}
}
