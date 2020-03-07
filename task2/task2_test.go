package task2

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 3",
			args: args{
				N: 12,
				S: "1A 2A,12A 12A",
				T: "12A",
			},
			want: "1,0",
		},
		{
			name: "case 1",
			args: args{
				N: 4,
				S: "1B 2C,2D 4D",
				T: "2B 2D 3D 4D 4A",
			},
			want: "1,1",
		},
		{
			name: "case 2",
			args: args{
				N: 3,
				S: "1A 1B,2C 2C",
				T: "1B",
			},
			want: "0,1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
