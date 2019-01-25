package Leetcode_543_DiameterOfBinaryTree

import (
	. "leetcode-golang/helper"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Case1",
			args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
			},
			1,
		},
		{
			"Case2",
			args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			2,
		},

		{
			"Case3",
			args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{Val: 3},
				},
			},
			3,
		},

		{
			"Case4",
			args{
				root: nil,
			},
			0,
		},
		{
			"Case5",
			args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 4,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					}},
			},
			4,
		},

		{
			"Case6",
			args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 5,
								Left: &TreeNode{
									Val: 7,
								},
							},
							Right: &TreeNode{
								Val: 4,
								Right: &TreeNode{
									Val: 6,
									Right: &TreeNode{
										Val: 8,
									},
								},
							},
						},
					},
					Right: &TreeNode{Val: 9},
				},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
