package leetcode617

import (
	. "Leetcode-golang/helper"
	"fmt"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"Case1",
			args{
				t1: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2}},
				t2: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2}},
			},

			&TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
		},
		{
			"Case2",
			args{
				t1: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2}},
				t2: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2}},
			},

			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 2},
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

func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
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
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				t1: &TreeNode{
					Val: 1,
				},
				t2: &TreeNode{
					Val: 1,
				},
			},
			true,
		},
		{
			"case2",
			args{
				t1: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
				t2: &TreeNode{
					Val: 1,
				},
			},
			false,
		},
		{
			"case3",
			args{
				t1: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
				t2: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 3},
				},
			},
			false,
		},
		{
			"case3",
			args{
				t1: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 2},
				},
				t2: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
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
