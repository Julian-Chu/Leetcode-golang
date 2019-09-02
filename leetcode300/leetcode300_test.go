package leetcode300

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[]",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "[2, 3, 7 ,1]",
			args: args{
				nums: []int{2, 3, 7, 1},
			},
			want: 3,
		},
		{
			name: "[10,9,2,5,3,7,101,18]",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "[4,10,4,3,8,9]",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
