package LeetCode_309_BestTimeToBuyAndSellStockWithCooldown

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	s0 := 0          //rest after rest, rest after sell
	s1 := -prices[0] // buy after rest, rest after buy
	s2 := -1 << 31   //  sell after rest

	for i := 0; i < len(prices); i++ {
		pre0 := s0
		pre1 := s1
		pre2 := s2

		s0 = max(pre0, pre2)
		s1 = max(pre0-prices[i], pre1)
		s2 = s1 + prices[i]
	}

	return max(s0, s2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit_1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])) //hold stock
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])                                      //sold stock, now in cooldown
		dp[i][2] = dp[i-1][0] + prices[i]                                           // sell stock today
		dp[i][3] = dp[i-1][2]                                                       // today is cool down
	}

	return max(dp[n-1][3], max(dp[n-1][1], dp[n-1][2]))
}
