package leetcode94

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "case2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{4, 1, 3, 2},
		},
		{
			name: "case3",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  1,
						Left: nil,
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
