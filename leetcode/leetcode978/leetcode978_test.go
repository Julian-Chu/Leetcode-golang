package leetcode978

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[9,4,2,10,7,8,8,1,9]",
			args: args{
				A: []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.A); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
