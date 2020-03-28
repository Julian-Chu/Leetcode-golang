package leetcode450

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[1,null,2] key=1",
			args: args{
				root: utils.Ints2Tree([]int{1, utils.NULL, 2}),
				key:  1,
			},
			want: &TreeNode{
				Val: 2,
			},
		},
		{
			name: "[5,3,6,2,4,null,7]",
			args: args{
				root: utils.Ints2Tree([]int{5, 3, 6, 2, 4, utils.NULL, 7}),
				key:  3,
			},
			want: utils.Ints2Tree([]int{5, 2, 6, utils.NULL, 4, utils.NULL, 7}),
		},
		{
			name: "[2,1,3]",
			args: args{
				root: utils.Ints2Tree([]int{2, 1, 3}),
				key:  2,
			},
			want: utils.Ints2Tree([]int{1, utils.NULL, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
