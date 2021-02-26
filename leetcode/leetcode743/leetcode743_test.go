package leetcode743

import "testing"

func Test_networkDelayTimeHeap(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
				N:     4,
				K:     2,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}},
				N:     3,
				K:     1,
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				times: [][]int{{1, 2, 1}, {2, 1, 3}},
				N:     2,
				K:     2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTimeHeap(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTimeHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
