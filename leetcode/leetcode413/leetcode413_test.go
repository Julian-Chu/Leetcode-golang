package leetcode413

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "[1,2]",
			args: args{
				A: []int{1, 2},
			},
			want: 0,
		},
		{
			name: "[1,2,3,4]",
			args: args{
				A: []int{1, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "[7,7,7,7]",
			args: args{
				A: []int{7, 7, 7, 7},
			},
			want: 3,
		},
		{
			name: "[1,1,2,5,7]",
			args: args{
				A: []int{1, 1, 2, 5, 7},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.A); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
