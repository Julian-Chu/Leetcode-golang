package leetcode475

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		house   []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[1,2,3],[2], return 1", args{[]int{1, 2, 3}, []int{2}}, 1},
		{"[1,2,3,4,5],[3], return 2", args{[]int{1, 2, 3, 4, 5}, []int{3}}, 2},
		{"[1,2,3,4],[1,4], return 1", args{[]int{1, 2, 3, 4}, []int{1, 4}}, 1},
		{"[1,2,3,5,15],[2,30], return 13", args{[]int{1, 2, 3, 5, 15}, []int{2, 30}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.house, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
