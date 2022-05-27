package leetcode23

import (
	"reflect"
	"testing"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case",
			args: args{
				lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name: "[]",
			args: args{
				lists: [][]int{},
			},
			want: []int{},
		},
		{
			name: "[[]]",
			args: args{
				lists: [][]int{{}},
			},
			want: []int{},
		},
		{
			name: "[[1]]",
			args: args{
				lists: [][]int{{1}},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.args.lists))
			for _, list := range tt.args.lists {
				lists = append(lists, utils.IntsToListNode(list))
			}
			if got := mergeKLists(lists); !reflect.DeepEqual(utils.ListNodeToInts(got), tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
