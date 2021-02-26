package leetcode680

func validPalindrome(s string) bool {
	left, right := findDiff(s, 0, len(s)-1)

	if left > right {
		return true
	}

	return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
}

func isPalindrome(s string, left, right int) bool {
	l, r := findDiff(s, left, right)
	return l >= r
}

func findDiff(s string, left, right int) (int, int) {
	for left < right {
		if s[left] != s[right] {
			break
		}
		left++
		right--
	}
	return left, right
}
