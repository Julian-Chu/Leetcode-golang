package leetcode179

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 10},
			},
			want: "210",
		},
		{
			name: "case2",
			args: args{
				nums: []int{34, 30},
			},
			want: "3430",
		},
		{
			name: "case3",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "9534330",
		},
		{
			name: "case4",
			args: args{
				nums: []int{34, 3},
			},
			want: "343",
		},
		{
			name: "case5",
			args: args{
				nums: []int{121, 12},
			},
			want: "12121",
		},
		{
			name: "case6",
			args: args{
				nums: []int{0, 0},
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
