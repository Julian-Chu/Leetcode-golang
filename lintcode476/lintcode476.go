package lintcode476

/**
 * @param A: An integer array
 * @return: An integer
 */
func stoneGame(A []int) int {
	if len(A) < 2 {
		return 0
	}

	n := len(A)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		for j := range dp[0] {
			dp[i][j] = 1<<31 - 1
			if i == j {
				dp[i][j] = 0
			}
		}
	}

	prefixSum := make([]int, n+1)
	for i := range prefixSum {
		if i == 0 {
			continue
		}
		prefixSum[i] = prefixSum[i-1] + A[i-1]
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < (n - l + 1); i++ {

			j := i + l - 1
			for k := i; k < j; k++ {
				sum := prefixSum[j+1] - prefixSum[i]
				dp[i][j] = min(dp[i][j], (dp[i][k] + dp[k+1][j] + sum))
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
