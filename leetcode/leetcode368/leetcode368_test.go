package leetcode368

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[]",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "[1]",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "[1,2,3]",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2},
		},
		{
			name: "[1,2,4,8]",
			args: args{
				nums: []int{1, 2, 4, 8},
			},
			want: []int{1, 2, 4, 8},
		},
		{
			name: "[1,2,3,6]",
			args: args{
				nums: []int{1, 2, 3, 6},
			},
			want: []int{1, 2, 6},
		},
		{
			name: "[1,2,4,8,16]",
			args: args{
				nums: []int{1, 2, 4, 8, 16},
			},
			want: []int{1, 2, 4, 8, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largetDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
