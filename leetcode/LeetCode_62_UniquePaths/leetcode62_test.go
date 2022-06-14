package LeetCode_62_UniquePaths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m=3,n=2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "m=23, n=12",
			args: args{
				m: 23,
				n: 12,
			},
			want: 193536720,
		},
		{
			name: "m=19, n=13",
			args: args{
				m: 19,
				n: 13,
			},
			want: 86493225,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
