package factorialtrailingzeroes

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"input 3 return 0", args{n: 3}, 0},
		{"input 5 return 1", args{n: 5}, 1},
		{"input 10 return 2", args{n: 10}, 2},
		{"input 30 return 7", args{n: 30}, 7},
		{"input 60 return 14", args{n: 60}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
