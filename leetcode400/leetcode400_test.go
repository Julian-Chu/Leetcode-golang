package leetcode400

import "testing"

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "13",
			args: args{
				n: 13,
			},
			want: 1,
		},
		{
			name: "10",
			args: args{
				n: 10,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "11",
			args: args{
				n: 11,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
