package leetcode526

import "testing"

func Test_countArrangement(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "N=1",
			args: args{
				N: 1,
			},
			want: 1,
		},
		{
			name: "N=2",
			args: args{
				N: 2,
			},
			want: 2,
		},
		{
			name: "N=3",
			args: args{
				N: 3,
			},
			want: 3,
		},
		{
			name: "N=4",
			args: args{
				N: 4,
			},
			want: 8,
		},
		{
			name: "N=5",
			args: args{
				N: 5,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrangement(tt.args.N); got != tt.want {
				t.Errorf("countArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
