package leetcode389

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "abcd abcde",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e',
		},
		{
			name: "a aa",
			args: args{
				s: "a",
				t: "aa",
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
