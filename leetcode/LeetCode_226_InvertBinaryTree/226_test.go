package LeetCode_226_InvertBinaryTree

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *utils.TreeNode
	}{
		{
			"Case 1",
			args{
				root: &utils.TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2},
				},
			},
			&utils.TreeNode{
				Val:   1,
				Right: &utils.TreeNode{Val: 2},
			},
		},
		{
			"Case 2",
			args{
				root: &utils.TreeNode{
					Val: 1,
					Left: &utils.TreeNode{
						Val:  2,
						Left: &utils.TreeNode{Val: 3},
					},
				},
			},
			&utils.TreeNode{
				Val:   1,
				Right: &utils.TreeNode{Val: 2, Right: &utils.TreeNode{Val: 3}},
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
