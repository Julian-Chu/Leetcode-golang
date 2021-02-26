package leetcode25

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
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
			want: []int{4, 3, 2, 1},
		},
		{
			name: "1->2",
			args: args{
				head: []int{1, 2},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		head := utils.IntsToListNode(tt.args.head)
		t.Run(tt.name, func(t *testing.T) {
			tail := head
			for tail.Next != nil {
				tail = tail.Next
			}
			if got, _ := reverse(head, tail); !reflect.DeepEqual(utils.ListNodeToInts(got), tt.want) {
				t.Errorf("reverse() = %v, want %v", utils.ListNodeToInts(got), tt.want)
			}
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1->2",
			args: args{
				head: []int{1, 2},
				k:    2,
			},
			want: []int{2, 1},
		},
		{
			name: "1->2->3->4->5,k=2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "1->2->3->4->5,k=3",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.IntsToListNode(tt.args.head)
			want := utils.IntsToListNode(tt.want)
			if got := reverseKGroup(head, tt.args.k); !reflect.DeepEqual(got, want) {
				t.Errorf("reverseKGroup() = %v, want %v", utils.ListNodeToInts(got), tt.want)
			}
		})
	}
}
