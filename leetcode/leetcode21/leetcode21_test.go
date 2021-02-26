package leetcode21

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				l1: utils.IntsToListNode([]int{1, 2, 4}),
				l2: utils.IntsToListNode([]int{1, 3, 4}),
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			ints := utils.ListNodeToInts(got)
			if !reflect.DeepEqual(ints, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", ints, tt.want)
			}
		})
	}
}
