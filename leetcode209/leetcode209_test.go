package leetcode209

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "s=7, nums=[2,3,1,2,4,3]",
			args: args{
				s:    7,
				nums: []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			name: "s=11, nums=[1,2,3,4,5]",
			args: args{
				s:    11,
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
