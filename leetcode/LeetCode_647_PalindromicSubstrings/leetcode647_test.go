package LeetCode_647_PalindromicSubstrings

import (
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{"a"}, 1},
		{"aa", args{"aa"}, 3},
		{"ab", args{"ab"}, 2},
		{"abc", args{"abc"}, 3},
		{"aaa", args{"aaa"}, 6},
		{"fdsklf", args{"fdsklf"}, 6},
		{"babab", args{"babab"}, 9},
		{"baabab", args{"baabab"}, 10},
		{"bbccaacacdbdbcbcbbbcbadcbdddbabaddbcadb", args{"bbccaacacdbdbcbcbbbcbadcbdddbabaddbcadb"}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
