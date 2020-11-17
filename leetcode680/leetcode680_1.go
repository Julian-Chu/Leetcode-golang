package leetcode680

func validPalindrome_1(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isPal(s, left+1, right) || isPal(s, left, right-1)
		}
		left++
		right--
	}

	return true
}

func isPal(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
