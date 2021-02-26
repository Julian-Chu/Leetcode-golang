package leetcode189

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums        []int
		k           int
		rotatedNums []int
	}
	var tests = []struct {
		name string
		args args
	}{
		{
			name: "[1,2,3,4,5,6,7]",
			args: args{
				nums:        []int{1, 2, 3, 4, 5, 6, 7},
				k:           3,
				rotatedNums: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			for i := range tt.args.rotatedNums {
				if tt.args.nums[i] != tt.args.rotatedNums[i] {
					t.Error("Not equal")
				}
			}

		})
	}
}
