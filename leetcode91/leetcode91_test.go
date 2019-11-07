package leetcode91

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "10",
			args: args{
				s: "10",
			},
			want: 1,
		},
		{
			name: "00",
			args: args{
				s: "00",
			},
			want: 0,
		},
		{
			name: "01",
			args: args{
				s: "01",
			},
			want: 0,
		},
		{
			name: "12",
			args: args{
				s: "12",
			},
			want: 2,
		},
		{
			name: "226",
			args: args{
				s: "226",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
