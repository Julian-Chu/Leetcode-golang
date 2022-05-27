package leetcode337

import (
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

var int2TreeNode = utils.Ints2Tree

func Test_rob(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,1,3,null,4]",
			args: args{
				root: []int{2, 1, 3, utils.NULL, 4},
			},
			want: 7,
		},
		{
			name: "[3,2,3,null,3,null,1]",
			args: args{
				root: []int{3, 2, 3, utils.NULL, 3, utils.NULL, 1},
			},
			want: 7,
		},
		{
			name: "[3,4,5,1,3,null,1]",
			args: args{
				root: []int{3, 4, 5, 1, 3, utils.NULL, 1},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(int2TreeNode(tt.args.root)); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
