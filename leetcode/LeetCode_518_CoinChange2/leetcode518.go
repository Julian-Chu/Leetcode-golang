package LeetCode_518_CoinChange2

func change(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}
