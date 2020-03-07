package leetcode202

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "19",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "100",
			args: args{
				n: 100,
			},
			want: true,
		},
		{
			name: "18",
			args: args{
				n: 18,
			},
			want: false,
		},
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
