package leetcode33

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,1], target = 1",
			args: args{
				nums:   []int{2, 3, 1},
				target: 1,
			},
			want: 2,
		},
		{
			name: "[2,3,1], target = 3",
			args: args{
				nums:   []int{2, 3, 1},
				target: 3,
			},
			want: 1,
		},
		{
			name: "[4,5,6,7,0,1,2], target = 0",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "[4,5,6,7,0,1,2], target = 0",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "[]",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: -1,
		},
		{
			name: "[1]",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "[1,3]",
			args: args{
				nums:   []int{1, 3},
				target: 3,
			},
			want: 1,
		},
		{
			name: "[1,2,3]",
			args: args{
				nums:   []int{1, 2, 3},
				target: 3,
			},
			want: 2,
		},
		{
			name: "[1,2,3]",
			args: args{
				nums:   []int{1, 2, 3},
				target: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
