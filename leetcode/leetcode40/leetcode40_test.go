package leetcode40

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
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
			name: "[1,2] 3",
			args: args{
				candidates: []int{1, 2},
				target:     3,
			},
			want: [][]int{{1, 2}},
		},
		{
			name: "[2,2,3], 6",
			args: args{
				candidates: []int{2, 2, 3},
				target:     6,
			},
			want: [][]int{},
		}, {
			name: "[1,1,2]  3",
			args: args{
				candidates: []int{1, 1, 2},
				target:     3,
			},
			want: [][]int{{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
