package leetcode168

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: "A",
		},
		{
			name: "28",
			args: args{
				n: 28,
			},
			want: "AB",
		},
		{
			name: "701",
			args: args{
				n: 701,
			},
			want: "ZY",
		},
		{
			name: "52",
			args: args{
				n: 52,
			},
			want: "AZ",
		},
		{
			name: "27",
			args: args{
				n: 27,
			},
			want: "AA",
		},
		{
			name: "703",
			args: args{
				n: 703,
			},
			want: "AAA",
		},
		{
			name: "1048",
			args: args{
				n: 1048,
			},
			want: "ANH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
