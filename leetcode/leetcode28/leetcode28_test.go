package leetcode28

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hello ll",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "aaaaa bba",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "aaaaa  ",
			args: args{
				haystack: "aaaaa",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "aaaa",
			args: args{
				haystack: "aaaa",
				needle:   "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
