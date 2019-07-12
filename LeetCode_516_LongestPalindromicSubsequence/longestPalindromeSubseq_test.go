package leetcode516longestpalindromicsubsequence

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"b", args{"b"}, 1},
		{"bb", args{"bb"}, 2},
		{"bbb", args{"bbb"}, 3},
		{"bbba", args{"bbba"}, 3},
		{"bbbab", args{"bbbab"}, 4},
		{"cbbd", args{"cbbd"}, 2},
		{"aabaa", args{"aabaa"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
