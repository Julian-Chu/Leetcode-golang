package leetcode279

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 1<<32 - 1
	}
	perfects := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
