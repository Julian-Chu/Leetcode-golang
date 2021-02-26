package leetcode372

import "testing"

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a=2,b=[3]",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "a=2,b=[1,2]",
			args: args{
				a: 2,
				b: []int{1, 2},
			},
			want: 85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
