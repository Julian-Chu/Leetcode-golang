package leetcode166

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1/2",
			args: args{
				numerator:   1,
				denominator: 2,
			},
			want: "0.5",
		},
		{
			name: "2/1",
			args: args{
				numerator:   2,
				denominator: 1,
			},
			want: "2",
		},
		{
			name: "2/3",
			args: args{
				numerator:   2,
				denominator: 3,
			},
			want: "0.(6)",
		},
		{
			name: "4/333",
			args: args{
				numerator:   4,
				denominator: 333,
			},
			want: "0.(012)",
		},
		{
			name: "1/6",
			args: args{
				numerator:   1,
				denominator: 6,
			},
			want: "0.1(6)",
		},
		{
			name: "-50/8",
			args: args{
				numerator:   -50,
				denominator: 8,
			},
			want: "-6.25",
		},
		{
			name: "2147483647",
			args: args{
				numerator:   2147483647,
				denominator: 1,
			},
			want: "2147483647",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
