package LeetCode_494_TargetSum

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1,1,1,1]",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				S:    3,
			},
			want: 5,
		},
		{
			name: "[1,0,1]",
			args: args{
				nums: []int{1, 0, 1},
				S:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
