package LeetCode_120_Triangle

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			min := triangle[i+1][j]
			if triangle[i+1][j] > triangle[i+1][j+1] {
				min = triangle[i+1][j+1]
			}
			triangle[i][j] += min
		}
	}

	return triangle[0][0]
}

func minimumTotal_1(triangle [][]int) int {
	n := len(triangle)

	// distance from (i, j) to bottom
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
