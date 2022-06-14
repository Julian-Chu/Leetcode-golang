package LeetCode_70_ClimbingStairs

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "6",
			args: args{
				n: 6,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
