package LeetCode_367_ValidPerfectSquare

import "testing"

func Test_isPerfectSquare(t *testing.T) {
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
			want: true,
		},
		{
			name: "1",
			args: args{
				num: 1,
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
		{
			name: "14",
			args: args{
				num: 14,
			},
			want: false,
		},
		{
			name: "16",
			args: args{
				num: 16,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
