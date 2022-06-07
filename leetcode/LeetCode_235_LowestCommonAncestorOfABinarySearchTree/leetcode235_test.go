package LeetCode_235_LowestCommonAncestorOfABinarySearchTree

import (
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		nodes []int
		p     []int
		q     []int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[6,2,8,0,4,7,9,null,null,3,5]",
			args: args{
				nodes: []int{6, 2, 8, 0, 4, 7, 9, utils.NULL, utils.NULL, 3, 5},
				p:     []int{2},
				q:     []int{8},
			},
			want: []int{6},
		},
		{
			name: "[6,2,8,0,4,7,9,null,null,3,5]",
			args: args{
				nodes: []int{6, 2, 8, 0, 4, 7, 9, utils.NULL, utils.NULL, 3, 5},
				p:     []int{2},
				q:     []int{4},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(utils.Ints2Tree(tt.args.nodes), utils.Ints2Tree(tt.args.p), utils.Ints2Tree(tt.args.q)); got.Val != tt.want[0] {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, utils.Ints2Tree(tt.want).Val)
			}
		})
	}
}
