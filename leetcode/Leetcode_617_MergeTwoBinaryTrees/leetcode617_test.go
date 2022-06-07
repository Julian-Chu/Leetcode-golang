package leetcode617

import (
	"fmt"
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *utils.TreeNode
		t2 *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *utils.TreeNode
	}{
		{
			"Case1",
			args{
				t1: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2}},
				t2: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2}},
			},

			&utils.TreeNode{
				Val:  2,
				Left: &utils.TreeNode{Val: 4},
			},
		},
		{
			"Case2",
			args{
				t1: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2}},
				t2: &utils.TreeNode{
					Val:   1,
					Right: &utils.TreeNode{Val: 2}},
			},

			&utils.TreeNode{
				Val:   2,
				Left:  &utils.TreeNode{Val: 2},
				Right: &utils.TreeNode{Val: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); !compareTree(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareTree(t1 *utils.TreeNode, t2 *utils.TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}

	if t2 == nil {
		fmt.Println(t1.Val)
		return false
	}

	if t1.Val != t2.Val {
		fmt.Println(t1.Val, " ", t2.Val)
		return false
	}
	return compareTree(t1.Left, t2.Left) && compareTree(t1.Right, t2.Right)

}

func xTest_compareTree(t *testing.T) {
	type args struct {
		t1 *utils.TreeNode
		t2 *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				t1: &utils.TreeNode{
					Val: 1,
				},
				t2: &utils.TreeNode{
					Val: 1,
				},
			},
			true,
		},
		{
			"case2",
			args{
				t1: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2},
				},
				t2: &utils.TreeNode{
					Val: 1,
				},
			},
			false,
		},
		{
			"case3",
			args{
				t1: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2},
				},
				t2: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 3},
				},
			},
			false,
		},
		{
			"case3",
			args{
				t1: &utils.TreeNode{
					Val:   1,
					Left:  &utils.TreeNode{Val: 2},
					Right: &utils.TreeNode{Val: 2},
				},
				t2: &utils.TreeNode{
					Val:   1,
					Left:  &utils.TreeNode{Val: 2},
					Right: &utils.TreeNode{Val: 3},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTree(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("compareTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
