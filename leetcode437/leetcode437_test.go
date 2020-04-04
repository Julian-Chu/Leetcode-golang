package leetcode437

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[0,1,1]",
			args: args{
				root: []int{0, 1, 1},
				sum:  1,
			},
			want: 4,
		},
		{
			name: "[1,-2,-3,1,3,-2,null,-1]",
			args: args{
				root: []int{1, -2, -3, 1, 3, -2, utils.NULL, -1},
				sum:  3,
			},
			want: 1,
		},
		{
			name: "[10,5,-3,3,2,null,11,3,-2,null,1]",
			args: args{
				root: []int{10, 5, -3, 3, 2, utils.NULL, 11, 3, -2, utils.NULL, 1},
				sum:  8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(utils.Ints2Tree(tt.args.root), tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
