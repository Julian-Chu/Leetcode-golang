package leetcode110

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case1",
			args{
				root: &TreeNode{Val: 2},
			},
			true,
		},

		{"Case2",
			args{
				root: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
			},
			true,
		},

		{"Case3",
			args{
				root: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
			},
			false,
		},

		{"Case4",
			args{
				root: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 3,
						Left: &TreeNode{Val: 4,
							Right: &TreeNode{Val: 6}}},
					Right: &TreeNode{Val: 10},
				},
			},
			false,
		},

		{"Case5",
			args{
				root: &TreeNode{Val: 3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{Val: 20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
