package leetcode350

import (
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,2,1] [2,2]",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2, 2},
		},
		{
			name: "[4,9,5],[9,4,9,8,4]",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.args.nums1, tt.args.nums2)
			if len(got) != len(tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
			m := make(map[int]int)
			for _, num := range got {
				m[num]++
			}

			for _, num := range tt.want {
				m[num]--
			}

			for _, v := range m {
				if v != 0 {
					t.Errorf("intersect() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
