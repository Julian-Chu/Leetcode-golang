package LeetCode_322_CoinChange

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1

		for _, c := range coins {
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func coinChange_1(coins []int, amount int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = 1<<31 - 1
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == 1<<31-1 {
		return -1
	}
	return dp[amount]
}
