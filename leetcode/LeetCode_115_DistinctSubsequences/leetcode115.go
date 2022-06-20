package LeetCode_115_DistinctSubsequences

func numDistinct(s string, t string) int {
	if s == "" {
		return 0
	}

	m := len(s)
	n := len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}
	return dp[m][n]
}
