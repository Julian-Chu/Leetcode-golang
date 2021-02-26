package leetcode107

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &utils.TreeNode{
						Val: 9,
					},
					Right: &utils.TreeNode{
						Val:   20,
						Left:  &utils.TreeNode{Val: 15},
						Right: &utils.TreeNode{Val: 7},
					},
				},
			},
			want: [][]int{{15, 7}, {9, 20}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
