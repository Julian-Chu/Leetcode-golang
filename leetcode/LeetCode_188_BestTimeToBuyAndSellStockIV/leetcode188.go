package LeetCode_188_BestTimeToBuyAndSellStockIV

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}

	for j := range dp[0] {
		if j&1 == 1 {
			dp[0][j] = -prices[0]
		}
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}

	return dp[len(prices)-1][k*2]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
