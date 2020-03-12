package leetcode319

import "testing"

func Test_bulbSwitch(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				n: 6,
			},
			want: 2,
		},
		{
			name: "8",
			args: args{
				n: 8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.args.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
