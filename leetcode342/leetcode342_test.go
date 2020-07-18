package leetcode342

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0",
			args: args{
				num: 0,
			},
			want: false,
		},
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: true,
		},
		{
			name: "16",
			args: args{
				num: 16,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				num: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
