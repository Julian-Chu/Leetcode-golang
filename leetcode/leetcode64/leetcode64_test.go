package leetcode64

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[]]",
			args: args{
				grid: [][]int{},
			},
			want: 0,
		},
		{
			name: "[[1]]",
			args: args{
				grid: [][]int{{1}},
			},
			want: 1,
		},
		{
			name: "[[1 ,2,3]]",
			args: args{
				grid: [][]int{{1, 2, 3}},
			},
			want: 6,
		},
		{
			name: "[[1,3,1],[1,5,1],[4,2,1]]",
			args: args{
				grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			},
			want: 7,
		},
		{
			name: "[1,2],[1,1]",
			args: args{
				grid: [][]int{{1, 2}, {1, 1}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
