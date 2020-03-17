package leetcode227

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2*2",
			args: args{
				s: "2*2",
			},
			want: 4,
		},
		{
			name: "11+12",
			args: args{
				s: "11+12",
			},
			want: 23,
		},
		{
			name: " 11 + 12 ",
			args: args{
				s: " 11 + 12 ",
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
