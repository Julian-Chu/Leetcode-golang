package LeetCode_54_SpiralMatrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[[1,2],[3,4]]",
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
			},
			want: []int{1, 2, 4, 3},
		},
		{
			name: "[[1,2,3],[4,5,6],[7,8,9]]",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "[[1,2,3],[4,5,6]]",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: []int{1, 2, 3, 6, 5, 4},
		},
		{
			name: "[[1,2,3,4],[10,11,12,13]]",
			args: args{
				matrix: [][]int{{1, 2, 3, 4}, {10, 11, 12, 13}},
			},
			want: []int{1, 2, 3, 4, 13, 12, 11, 10},
		},
		{
			name: "[[1,2],[3,4],[5,6],[7,8]]",
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			},
			want: []int{1, 2, 4, 6, 8, 7, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
