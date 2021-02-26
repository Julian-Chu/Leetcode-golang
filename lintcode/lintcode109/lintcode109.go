package lintcode109

func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return -1
	}

	n := len(triangle)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	minPath := 1<<32 - 1

	for j := 0; j < n; j++ {
		if dp[n-1][j] < minPath {
			minPath = dp[n-1][j]
		}
	}
	return minPath
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
