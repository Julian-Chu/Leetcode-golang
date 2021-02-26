package leetcode213

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,2]",
			args: args{
				nums: []int{2, 3, 2},
			},
			want: 3,
		},
		{
			name: "[1,2,3,1]",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "[1,0,2]",
			args: args{
				nums: []int{1, 0, 2},
			},
			want: 2,
		},
		{
			name: "[1,0]",
			args: args{
				nums: []int{1, 0},
			},
			want: 1,
		},
		{
			name: "[]",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "[1]",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
