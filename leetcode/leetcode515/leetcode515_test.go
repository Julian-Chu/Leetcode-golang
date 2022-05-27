package leetcode515

import (
	"reflect"
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_largestValues(t *testing.T) {
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
				root: utils.Ints2Tree([]int{1, 3, 2, 5, 3, utils.NULL, 9}),
			},
			want: []int{1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
