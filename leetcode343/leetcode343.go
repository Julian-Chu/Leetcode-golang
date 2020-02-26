package leetcode343

func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}

	dp := make([]int, n+1)

	dp[2] = 2
	dp[3] = 3

	for i := 4; i < n+1; i++ {
		max := 0
		for j := 1; j < i; j++ {
			res := j * dp[i-j]
			if res > max {
				max = res
			}
		}
		dp[i] = max
	}

	return dp[n]
}
