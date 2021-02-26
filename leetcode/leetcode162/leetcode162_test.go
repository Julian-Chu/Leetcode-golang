package leetcode162

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,1]",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: []int{2},
		},
		{
			name: "[1,2,1,3,5,6,4]",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: []int{1, 5},
		},
		{
			name: "[1,2]",
			args: args{
				nums: []int{1, 2},
			},
			want: []int{1},
		},
		{
			name: "[2,1]",
			args: args{
				nums: []int{2, 1},
			},
			want: []int{0},
		},
		{
			name: "[1,3,2,1]",
			args: args{
				nums: []int{1, 3, 2, 1},
			},
			want: []int{1},
		},
	}

	var contains = func(answers []int, got int) bool {
		for _, v := range answers {
			if got == v {
				return true
			}
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); !contains(tt.want, got) {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
