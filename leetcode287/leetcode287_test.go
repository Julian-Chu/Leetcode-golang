package leetcode287

import "testing"

func TestFindDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,3,4,2,2]",
			args: args{
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
