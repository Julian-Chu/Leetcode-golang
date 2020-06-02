package leetcode457

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[-2,-3,-9]",
			args: args{
				nums: []int{-2, -3, -9},
			},
			want: false,
		},
		{
			name: "1,1",
			args: args{
				nums: []int{1, 1},
			},
			want: true,
		},
		{
			name: "[3,1,2]",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: true,
		},
		{
			name: "[1,2,3,4,5]",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "[2,-1,1,-2,-2]",
			args: args{
				nums: []int{2, -1, 1, -2, -2},
			},
			want: false,
		},
		{
			name: "[2,-1,1,2,2]",
			args: args{
				nums: []int{2, -1, 1, 2, 2},
			},
			want: true,
		},
		{
			name: "[-1,2]",
			args: args{
				nums: []int{-1, 2},
			},
			want: false,
		},
		{
			name: "[-2,1,-1,-2,-2]",
			args: args{
				nums: []int{-2, 1, -1, -2, -2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
