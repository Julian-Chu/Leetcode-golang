package leetcode81

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[]",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: false,
		},
		{
			name: "[1,2,3]",
			args: args{
				nums:   []int{1, 2, 3},
				target: 1,
			},
			want: true,
		},
		{
			name: "[3,1,2]",
			args: args{
				nums:   []int{3, 1, 2},
				target: 1,
			},
			want: true,
		},
		{
			name: "[2,5,6,0,0,1,2]",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "[2,5,6,0,0,1,2]",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "[2,5,6,0,0,1,2]",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
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
