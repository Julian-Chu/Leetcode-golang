package leetcode396

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A = [4, 3, 2, 6]",
			args: args{
				A: []int{4, 3, 2, 6},
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRotateFunction(tt.args.A); got != tt.want {
				t.Errorf("maxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
