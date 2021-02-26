package leetcode397

import "testing"

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "8",
			args: args{
				n: 8,
			},
			want: 3,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 4,
		},
		{
			name: "9",
			args: args{
				n: 9,
			},
			want: 4,
		},
		{
			name: "65535",
			args: args{
				n: 65535,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
