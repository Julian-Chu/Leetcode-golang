package leetcode347

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[4,1,-1,-1,2,3], k=2",
			args: args{
				nums: []int{4, 1, -1, 2, -1, 2, 3},
				k:    2,
			},
			want: []int{-1, 2},
		},
		{
			name: "[5,3,1,1,1,3,73,1], k =1",
			args: args{
				nums: []int{5, 3, 1, 1, 1, 3, 73, 1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "[3,0,1,0], k=1",
			args: args{
				nums: []int{3, 0, 1, 0},
				k:    1,
			},
			want: []int{0},
		},
		{
			name: "[1,1,1,2,2,3], k = 2",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "[1], k =1",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
