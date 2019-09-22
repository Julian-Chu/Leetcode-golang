package leetcode78

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[]",
			args: args{
				nums: []int{},
			},
			want: [][]int{{}},
		},
		{
			name: "[1]",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{}, {1}},
		},
		{
			name: "[1,2]",
			args: args{
				nums: []int{1, 2},
			},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
