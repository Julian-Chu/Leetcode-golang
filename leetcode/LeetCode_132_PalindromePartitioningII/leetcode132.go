package LeetCode_132_PalindromePartitioningII

func minCut(s string) int {
	n := len(s)
	isPalindromic := make([][]bool, n)
	for i := range isPalindromic {
		isPalindromic[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || isPalindromic[i+1][j-1]) {
				isPalindromic[i][j] = true
			}
		}
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = i //
	}

	for i := 1; i < n; i++ {
		if isPalindromic[0][i] {
			dp[i] = 0
			continue
		}

		for j := 0; j < i; j++ {
			if isPalindromic[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
