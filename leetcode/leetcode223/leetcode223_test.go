package leetcode223

import "testing"

func Test_computeArea(t *testing.T) {
	type args struct {
		A int
		B int
		C int
		D int
		E int
		F int
		G int
		H int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				A: -3,
				B: 0,
				C: 3,
				D: 4,
				E: 0,
				F: -1,
				G: 9,
				H: 2,
			},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeArea(tt.args.A, tt.args.B, tt.args.C, tt.args.D, tt.args.E, tt.args.F, tt.args.G, tt.args.H); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
