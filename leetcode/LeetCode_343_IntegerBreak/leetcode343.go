package LeetCode_343_IntegerBreak

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	res := 1

	for n > 4 {
		res *= 3
		n -= 3
	}

	return res * n
}

func integerBreak_dp(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
