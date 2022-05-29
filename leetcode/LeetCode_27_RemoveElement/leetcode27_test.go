package LeetCode_27_RemoveElement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,2,2,3]",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "[0,1,2,2,3,0,4,2]",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "[3,3]",
			args: args{
				nums: []int{3, 3},
				val:  3,
			},
			want: 0,
		},
		{
			name: "[4,5]",
			args: args{
				nums: []int{4, 5},
				val:  4,
			},
			want: 1,
		},

		{
			name: "[2]",
			args: args{
				nums: []int{2},
				val:  3,
			},
			want: 1,
		},
		{
			name: "[3,3]",
			args: args{
				nums: []int{3, 3},
				val:  5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
