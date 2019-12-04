package leetcode103

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
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
					Val: 1,
					Left: &utils.TreeNode{
						Val: 2,
						Left: &utils.TreeNode{
							Val: 4,
						},
					},
					Right: &utils.TreeNode{
						Val: 3,
						Right: &utils.TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
