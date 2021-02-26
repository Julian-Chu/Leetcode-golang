package leetcode230

import (
	"Leetcode-golang/utils"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,1,4,null,2] k=1",
			args: args{
				root: []int{3, 1, 4, utils.NULL, 2},
				k:    1,
			},
			want: 1,
		},
		{
			name: "[5,3,6,2,4,null,null,1]",
			args: args{
				root: []int{5, 3, 6, 2, 4, utils.NULL, utils.NULL, 1},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			root := utils.Ints2Tree(tt.args.root)
			if got := kthSmallest(root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
