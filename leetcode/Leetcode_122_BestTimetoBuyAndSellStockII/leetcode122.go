package leetcode122

func maxProfit(prices []int) int {
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func maxProfit_greedy(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := 0

	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}

	return profit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit_dp(prices []int) int {
	n := len(prices)

	dp := [2][]int{}
	// profit with holding stock
	dp[0] = make([]int, n)
	// profit without holding stock
	dp[1] = make([]int, n)

	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1]-prices[i])
		dp[1][i] = max(dp[1][i-1], dp[0][i-1]+prices[i])
	}

	return max(dp[0][n-1], dp[1][n-1])
}
