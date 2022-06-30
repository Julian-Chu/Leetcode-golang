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

func longestPalindrome_dp(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	substr := s[:1]
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			if dp[i][j] && j-i+1 > len(substr) {
				substr = s[i : j+1]
			}
		}
	}

	return substr
}
