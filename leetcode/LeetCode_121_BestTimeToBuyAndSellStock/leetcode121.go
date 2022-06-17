package LeetCode_121_BestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	max := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}
	return max
}

func maxProfit_1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])           //hold stock
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0]) // not hold stock
	}

	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
