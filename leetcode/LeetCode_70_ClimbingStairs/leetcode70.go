package LeetCode_70_ClimbingStairs

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	x, y := 1, 1
	for i := 2; i <= n; i++ {
		x, y = x+y, x
	}
	return x
}

func climbStairs_1(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs_dp(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}

	return dp[n]
}
