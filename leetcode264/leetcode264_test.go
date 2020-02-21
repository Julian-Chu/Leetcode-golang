package leetcode264

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "n=2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "n=10",
			args: args{
				n: 10,
			},
			want: 12,
		},
		{
			name: "n=1352",
			args: args{
				n: 1352,
			},
			want: 402653184,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
