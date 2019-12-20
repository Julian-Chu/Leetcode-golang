package leetcode144

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1, null,2,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &utils.TreeNode{Val: 2,
						Left: &TreeNode{Val: 3}},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "[1,2,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &utils.TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
