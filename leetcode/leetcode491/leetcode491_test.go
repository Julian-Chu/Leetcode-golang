package leetcode491

import (
	"reflect"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{
		//	name: "[1,2]",
		//	args: args{
		//		nums: []int{1, 2},
		//	},
		//	want: [][]int{{1, 2}},
		//},
		//{
		//	name: "[2,1,3]",
		//	args: args{
		//		nums: []int{2, 1, 3},
		//	},
		//	want: [][]int{{2, 3}, {1, 3}},
		//},
		{
			name: "[1,2,3]",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: [][]int{{2, 3}, {1, 3}, {1, 2, 3}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
