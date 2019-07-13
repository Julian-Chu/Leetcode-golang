package LeetCode_125_ValidPalindrome

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	var r []rune
	for _, v := range []rune(strings.ToLower(s)) {
		if isAlphbanumeric(v) {
			r = append(r, v)
		}
	}
	if r == nil {
		return true
	}

	fmt.Println(r)
	l := len(r)
	mid := l / 2
	b, e := mid, mid
	if l%2 == 0 {
		b, e = mid-1, mid
	}

	for b >= 0 && e < l {
		if r[b] != r[e] {
			return false
		}
		b--
		e++
	}
	return true
}

func isAlphbanumeric(r rune) bool {
	if 48 <= r && r <= 57 || 65 <= r && r <= 90 || 97 <= r && r <= 122 {
		return true
	}
	return false
}
