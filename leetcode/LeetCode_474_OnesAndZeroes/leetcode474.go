package LeetCode_474_OnesAndZeroes

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zero, one := 0, 0
		for j := range str {
			if str[j] == '0' {
				zero++
				continue
			}
			one++
		}

		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
