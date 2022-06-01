package LeetCode_24_SwapNodesInPairs

import (
	"reflect"
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1->2->3->4",
			args: args{
				head: []int{1, 2, 3, 4},
			},
			want: []int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.IntsToListNode(tt.args.head)
			want := utils.IntsToListNode(tt.want)
			if got := swapPairs(head); !reflect.DeepEqual(got, want) {
				t.Errorf("swapPairs() = %v, want %v", utils.ListNodeToInts(got), tt.want)
			}
		})
	}
}
