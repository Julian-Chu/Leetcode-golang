package leetcode203

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1->2->6->3->4->5->6",
			args: args{
				head: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "nil",
			args: args{
				head: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "1, v=1",
			args: args{
				head: []int{1},
				val:  1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.IntsToListNode(tt.args.head)
			want := utils.IntsToListNode(tt.want)
			if got := removeElements(head, tt.args.val); !reflect.DeepEqual(got, want) {
				t.Errorf("removeElements() = %v, want %v", utils.ListNodeToInts(got), want)
			}
		})
	}
}
