package leetcode404

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,3,4,5]",
			args: args{
				root: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "[3,9,20,null,null,15,7]",
			args: args{
				root: []int{3, 9, 20, utils.NULL, utils.NULL, 15, 7},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(utils.Ints2Tree(tt.args.root)); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
