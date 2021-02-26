package leetcode908

import "testing"

func Test_smallestRangeI(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A=[1],K=0", args{[]int{1}, 0}, 0},
		{"A=[0,10],K=2", args{[]int{0, 10}, 2}, 6},
		{"A=[2,7,2],K=1", args{[]int{2, 7, 2}, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeI(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("smallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
