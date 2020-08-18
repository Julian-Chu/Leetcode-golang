package leetcode328

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			name: "case 2",
			args: args{
				head: []int{2, 1, 3, 5, 6, 4, 7},
			},
			want: []int{2, 3, 6, 7, 1, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.IntsToListNode(tt.args.head)
			if got := oddEvenList(head); !reflect.DeepEqual(utils.ListNodeToInts(got), tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", utils.ListNodeToInts(got), tt.want)
			}
		})
	}
}
