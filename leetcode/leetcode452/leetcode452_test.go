package leetcode452

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[1,6],[2,8],[7,12],[10,16]]",
			args: args{
				points: [][]int{{1, 6}, {2, 8}, {7, 12}, {10, 16}},
			},
			want: 2,
		},
		{
			name: "[[1,6],[2,5], [6,7]]",
			args: args{
				points: [][]int{{1, 6}, {2, 5}, {6, 7}},
			},
			want: 2,
		},
		{
			name: "[[1,2],[2,3],[3,4],[4,5]]",
			args: args{
				points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
