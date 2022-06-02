package LeetCode_15_3Sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[-1,0,1,2,-1,-4]",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "[0,0,0,0]",
			args: args{
				nums: []int{0, 0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "[-2,0,0,2,2]",
			args: args{
				nums: []int{-2, 0, 0, 2, 2},
			},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "[1,1,-2]",
			args: args{
				nums: []int{1, 1, -2},
			},
			want: [][]int{{-2, 1, 1}},
		},
		{
			name: "[-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]",
			args: args{
				nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
