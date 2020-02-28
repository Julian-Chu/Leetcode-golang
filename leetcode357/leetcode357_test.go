package leetcode357

import "testing"

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 10,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 91,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 739,
		},
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
