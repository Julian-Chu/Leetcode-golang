package leetcode112

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &utils.TreeNode{Val: 2},
				},
				sum: 1,
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				root: &TreeNode{
					Val:   -2,
					Right: &utils.TreeNode{Val: -3},
				},
				sum: -5,
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				root: &TreeNode{
					Val:   0,
					Left:  &utils.TreeNode{Val: 1},
					Right: &utils.TreeNode{Val: 1},
				},
				sum: 0,
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &utils.TreeNode{
						Val: -2,
						Left: &TreeNode{
							Val:  1,
							Left: &TreeNode{Val: -1},
						},
						Right: &utils.TreeNode{Val: 3},
					},
					Right: &utils.TreeNode{Val: -3,
						Left: &utils.TreeNode{Val: -2}},
				},
				sum: -1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
