package LeetCode_18_4Sum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,0,-1,0,-2,2]",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "[0,0,0,0]",
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			want: [][]int{{0, 0, 0, 0}},
		},
		{
			name: "[-1,0,1,2,-1,-4]",
			args: args{
				nums:   []int{-1, 0, 1, 2, -1, -4},
				target: -1,
			},
			want: [][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
