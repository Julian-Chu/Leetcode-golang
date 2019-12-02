package utils

import "testing"

func Test_compareTrees(t *testing.T) {
	type args struct {
		node1 *TreeNode
		node2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				node1: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				node2: &TreeNode{},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				node1: &TreeNode{
					Val:   0,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
				node2: &TreeNode{
					Val:   0,
					Left:  &TreeNode{},
					Right: &TreeNode{},
				},
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				node1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{},
				},
				node2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{},
				},
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				node1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{},
				},
				node2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{},
				},
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				node1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{},
				},
				node2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTrees(tt.args.node1, tt.args.node2); got != tt.want {
				t.Errorf("compareTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
