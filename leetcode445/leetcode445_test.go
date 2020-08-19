package leetcode445

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				l1: []int{7, 2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 8, 0, 7},
		},
		{
			name: "case2",
			args: args{
				l1: []int{5},
				l2: []int{5},
			},
			want: []int{1, 0},
		},
		{
			name: "case3",
			args: args{
				l1: []int{9},
				l2: []int{8},
			},
			want: []int{1, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := utils.IntsToListNode(tt.args.l1)
			l2 := utils.IntsToListNode(tt.args.l2)
			if got := addTwoNumbers(l1, l2); !reflect.DeepEqual(utils.ListNodeToInts(got), tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", utils.ListNodeToInts(got), tt.want)
			}
		})
	}
}
