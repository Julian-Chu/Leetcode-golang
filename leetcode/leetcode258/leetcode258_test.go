package leetcode258

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "38",
			args: args{
				num: 38,
			},
			want: 2,
		},
		{
			name: "9",
			args: args{
				num: 9,
			},
			want: 9,
		},
		{
			name: "18",
			args: args{
				num: 18,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
