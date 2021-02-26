package leetcode456

import "testing"

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[-2,1,2,-2,1,2]",
			args: args{
				nums: []int{-2, 1, 2, -2, 1, 2},
			},
			want: true,
		},
		{
			name: "[1,2,3,4]",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "[3,1,4,2]",
			args: args{
				nums: []int{3, 1, 4, 2},
			},
			want: true,
		},
		{
			name: "[-1,3,2,0]",
			args: args{
				nums: []int{-1, 3, 2, 0},
			},
			want: true,
		},
		{
			name: "[3,4,0,1,2]",
			args: args{
				nums: []int{3, 4, 0, 1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
