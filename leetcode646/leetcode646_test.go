package leetcode646

import "testing"

func Test_findLongestChain(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[1,2], [2,3], [3,4]]",
			args: args{
				pairs: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			want: 2,
		},
		{
			name: "[[1,2], [5,6], [3,4]]",
			args: args{
				pairs: [][]int{{1, 2}, {5, 6}, {3, 4}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestChain(tt.args.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
