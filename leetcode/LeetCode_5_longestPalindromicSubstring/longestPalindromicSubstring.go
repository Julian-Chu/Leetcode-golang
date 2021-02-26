package LeetCode_5_longestPalindromicSubstring

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		first, last := i, i
		for last < len(s)-1 && s[last+1] == s[last] {
			last++
		}

		i = last + 1

		for last < len(s)-1 && first > 0 && s[last+1] == s[first-1] {
			last++
			first--
		}

		newLen := last + 1 - first
		if newLen > maxLen {
			start = first
			maxLen = newLen
		}

	}
	return s[start : start+maxLen]
}
