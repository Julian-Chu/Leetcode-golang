package leetcode92

import (
	"Leetcode-golang/utils"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		ints []int
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				ints: []int{1, 2, 3, 4, 5},
				m:    2,
				n:    4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "case2",
			args: args{
				ints: []int{3, 5},
				m:    1,
				n:    2,
			},
			want: []int{5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := utils.IntsToListNode(tt.args.ints)
			got := reverseBetween(node, tt.args.m, tt.args.n)
			res := utils.ListNodeToInts(got)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", res, tt.want)
			}
		})
	}
}
