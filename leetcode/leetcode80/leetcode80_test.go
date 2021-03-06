package leetcode80

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1,1,2,2,3]",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "[0,0,1,1,1,1,2,3,3]",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
		{
			name: "[1,1,1]",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
		{
			name: "[1,1,1,2,2,2,2]",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 2},
			},
			want: 4,
		},
		{
			name: "[1,1,1,1]",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 2,
		},
		{
			name: "[-3,-1,-1,0,0,0,0,0]",
			args: args{
				nums: []int{-3, -1, -1, 0, 0, 0, 0, 0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
