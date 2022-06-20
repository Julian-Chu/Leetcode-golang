package leetcode516longestpalindromicsubsequence

func longestPalindromeSubseq(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				continue
			}

			dp[i][j] = max(dp[i+1][j], dp[i][j-1])
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
