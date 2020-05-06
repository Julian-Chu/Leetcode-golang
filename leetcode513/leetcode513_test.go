package leetcode513

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_findBottomLeftValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: utils.Ints2Tree([]int{2, 1, 3}),
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				root: utils.Ints2Tree([]int{1, 2, 3, 4, utils.NULL, 5, 6, utils.NULL, utils.NULL, 7}),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
