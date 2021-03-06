package LeetCode_125_ValidPalindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphbanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphbanumeric(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphbanumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !(unicode.IsLetter(rune(s[left])) || unicode.IsDigit(rune(s[left]))) {
			left++
		}
		for left < right && !(unicode.IsLetter(rune(s[right])) || unicode.IsDigit(rune(s[right]))) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
