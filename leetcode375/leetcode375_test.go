package leetcode375

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: 3 + 1,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 4 + 2,
		},
		{
			name: "6",
			args: args{
				n: 6,
			},
			want: 3 + 5,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 4 + 6,
		},
		{
			name: "20",
			args: args{
				n: 20,
			},
			want: 49,
		},
		{
			name: "30",
			args: args{
				n: 30,
			},
			want: 79,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
