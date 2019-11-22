package leetcode171

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{
				s: "A",
			},
			want: 1,
		},
		{
			name: "AB",
			args: args{
				s: "AB",
			},
			want: 28,
		},
		{
			name: "ZY",
			args: args{
				s: "ZY",
			},
			want: 701,
		},
		{
			name: "AAA",
			args: args{
				s: "AAA",
			},
			want: 703,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
