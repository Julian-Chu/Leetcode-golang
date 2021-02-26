package leetcode415

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0 0",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "99, 12",
			args: args{
				num1: "99",
				num2: "12",
			},
			want: "111",
		},
		{
			name: "1 1",
			args: args{
				num1: "1",
				num2: "1",
			},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
