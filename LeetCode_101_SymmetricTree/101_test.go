package LeetCode_101_SymmetricTree

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1",
			args{
				&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			},
			false,
		},
		{"Case 2",
			args{
				&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
			},
			true,
		},
		{"Case 3",
			args{
				&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			},
			true,
		},
		{"Case 4",
			args{
				&TreeNode{Val: 1},
			},
			true,
		},

		{"Case 5",
			args{
				&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			},
			false,
		},

		{"Case 6",
			args{
				&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
