package LeetCode_435__NonOverlappingIntervals

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[1,2],[2,3],[3,4],[-100,-2],[5,7]]",
			args: args{
				intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {-100, -2}, {5, 7}},
			},
			want: 0,
		},
		{
			name: "[[1,2],[2,3],[3,4],[1,3]]",
			args: args{
				intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			},
			want: 1,
		},
		{
			name: "[[1,2],[1,2],[1,2]]",
			args: args{
				intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			},
			want: 2,
		},
		{
			name: "[[1,2],[2,3]]",
			args: args{
				intervals: [][]int{{1, 2}, {2, 3}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
