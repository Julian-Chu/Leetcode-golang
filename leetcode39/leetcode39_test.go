package leetcode39

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[2],2",
			args: args{
				candidates: []int{2},
				target:     2,
			},
			want: [][]int{{2}},
		},
		{
			name: "[1,2],2",
			args: args{
				candidates: []int{1, 2},
				target:     2,
			},
			want: [][]int{{1, 1}, {2}},
		},
		{
			name: "[2,3,6,7]",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		}, {
			name: "[2,3,5],8",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "[8,7,4,3],11",
			args: args{
				candidates: []int{8, 7, 4, 3},
				target:     11,
			},
			want: [][]int{{3, 4, 4}, {3, 8}, {4, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
