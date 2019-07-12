package leetcode516longestpalindromicsubsequence

func longestPalindromeSubseq(s string) int {
	i, j := 0, len(s)-1

	maxLen := 0

	if len(s) < 2 {
		return 1
	}

	for i <= j {
		if s[i] == s[j] {
			if i == j {
				maxLen++
				break
			}
			maxLen += 2
			i, j = i+1, j-1
		} else {
			left := longestPalindromeSubseq(s[i:j])
			right := longestPalindromeSubseq(s[i+1 : j+1])
			if left > right {
				maxLen += left
			} else {
				maxLen += right
			}
			return maxLen
		}
	}
	return maxLen
}
