package leetcode547

import "testing"

func Test_findCircleNum(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case3",
			args: args{
				M: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			},
			want: 1,
		},
		{
			name: "case1",
			args: args{
				M: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				M: [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
