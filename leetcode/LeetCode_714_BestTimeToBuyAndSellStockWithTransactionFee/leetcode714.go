package LeetCode_714_BestTimeToBuyAndSellStockWithTransactionFee

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	res := 0
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		if prices[i] >= minPrice && prices[i] <= minPrice+fee {
			continue
		}

		if prices[i] > minPrice+fee {
			res += prices[i] - minPrice - fee
			minPrice = prices[i] - fee // key step
		}
	}
	return res
}

func maxProfit_dp(prices []int, fee int) int {
	n := len(prices)
	dp := [2][]int{}
	dp[0] = make([]int, n) // with stock
	dp[1] = make([]int, n) // without stock

	dp[0][0] -= prices[0]
	for i := 1; i < n; i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1]-prices[i])
		dp[1][i] = max(dp[1][i-1], dp[0][i-1]+prices[i]-fee)
	}

	return max(dp[0][n-1], dp[1][n-1])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
