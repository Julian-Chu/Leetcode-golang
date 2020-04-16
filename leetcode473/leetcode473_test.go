package leetcode473

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,1,2,2,2]",
			args: args{
				nums: []int{1, 1, 2, 2, 2},
			},
			want: true,
		},
		{
			name: "[3,3,3,3,4]",
			args: args{
				nums: []int{3, 3, 3, 3, 4},
			},
			want: false,
		},
		{
			name: "[5,5,5,5,4,4,4,4,3,3,3,3]",
			args: args{
				nums: []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.nums); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
