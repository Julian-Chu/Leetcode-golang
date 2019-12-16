package leetcode129

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &utils.TreeNode{
						Val: 2,
					},
					Right: &utils.TreeNode{
						Val: 3,
					},
				},
			},
			want: 25,
		},
		{
			name: "[4,9,0,5,1]",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &utils.TreeNode{
						Val:   9,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 1},
					},

					Right: &utils.TreeNode{Val: 0},
				},
			},
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
