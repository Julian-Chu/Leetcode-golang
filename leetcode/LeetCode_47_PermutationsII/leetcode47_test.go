package LeetCode_47_PermutationsII

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,1,1]",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: [][]int{{1, 1, 1}},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 2},
			},
			want: [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1, 3},
			},
			want: [][]int{{1, 1, 1, 3}, {1, 1, 3, 1}, {1, 3, 1, 1}, {3, 1, 1, 1}},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1, 3, 3},
			},
			want: [][]int{{1, 1, 1, 3, 3}, {1, 1, 3, 1, 3}, {1, 3, 1, 1, 3}, {1, 1, 3, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
