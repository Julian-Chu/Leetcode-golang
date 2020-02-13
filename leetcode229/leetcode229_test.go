package leetcode229

import (
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[3,2,3]",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: []int{3},
		},
		{
			name: "[1,1,1,3,3,2,2,2]",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.args.nums)

			for _, v := range got {
				found := false
				for _, w := range tt.want {
					if v == w {
						found = true
					}
				}

				if !found {
					t.Errorf("majorityElement() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
