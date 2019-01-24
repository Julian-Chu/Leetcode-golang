package LeetCode_226_InvertBinaryTree

import (
	. "Leetcode-golang/helper"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"Case 1",
			args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
			},
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			"Case 2",
			args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
