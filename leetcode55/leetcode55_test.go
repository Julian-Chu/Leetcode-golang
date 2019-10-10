package leetcode55

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,0,4]",
			args: args{
				nums: []int{1, 0, 4},
			},
			want: false,
		},
		{
			name: "[1,1,4]",
			args: args{
				nums: []int{1, 1, 4},
			},
			want: true,
		},
		{
			name: "[2,3,1,1,4]",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "[3,2,1,0,4]",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "[0]",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
