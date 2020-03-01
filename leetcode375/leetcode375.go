package leetcode375

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 2; i < n+1; i++ {
		for j := i - 1; j > 0; j-- {
			if j+1 == i {
				dp[j][i] = j
				continue
			}
			min := 1 << 31

			for k := j + 1; k < i; k++ {
				max := dp[k+1][i]
				if dp[j][k-1] > max {
					max = dp[j][k-1]
				}
				max += k
				if min > max {
					min = max
				}
			}

			dp[j][i] = min
		}
	}
	return dp[1][n]
}
