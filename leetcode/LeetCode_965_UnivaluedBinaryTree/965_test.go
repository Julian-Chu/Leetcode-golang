package leetcode965

import "testing"

func Test_isUnivalTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case1",
			args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
			false,
		},
		{"Case2",
			args{&TreeNode{Val: 1, Left: &TreeNode{Val: 1}}},
			true,
		},
		{"Case3",
			args{&TreeNode{Val: 1, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
