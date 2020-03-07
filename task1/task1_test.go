package task1

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,1,8,2,4,1,1",
			args: args{
				A: []int{1, 1, 8, 2, 4, 1, 1},
			},
			want: true,
		},
		{
			name: "[1,1,1]",
			args: args{
				A: []int{1, 1, 1},
			},
			want: false,
		},
		{
			name: "[",
			args: args{
				A: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
