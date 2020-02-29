package leetcode375

func getMoneyAmount(n int) int {
	var dp []int
	if n < 3 {
		dp = make([]int, 3+1)
	} else {
		dp = make([]int, n+1)
	}
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	dp[3] = 2

	for i := 4; i < n+1; i++ {
		if i-1 > dp[i-4] {
			dp[i] = i - 3 + i - 1
		} else {
			dp[i] = i - 3 + dp[i-4]
		}
	}
	return dp[n]
}
