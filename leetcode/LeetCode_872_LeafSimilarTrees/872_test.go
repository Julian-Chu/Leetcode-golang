package leetcode872

import (
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Case1",
			args{
				root1: &TreeNode{
					Val: 1,
				},
				root2: &TreeNode{
					Val: 1,
				},
			},
			true,
		},

		{
			"Case2",
			args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				root2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			true,
		},

		{
			"Case3",
			args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				root2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
