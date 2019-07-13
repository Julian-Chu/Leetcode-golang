package reverseinteger

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"input 123 output 321", args{x: 123}, 321},
		{"input -123 output -321", args{x: -123}, -321},
		{"input 120 output 21", args{x: 120}, 21},
		{"input 1534236469 output 0", args{x: 1534236469}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
