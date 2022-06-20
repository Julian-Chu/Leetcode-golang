package leetcode583

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	res := make([][]int, n1+1)
	for i := range res {
		res[i] = make([]int, n2+1) // characters not need to change
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			res[i][j] = max(res[i-1][j], res[i][j-1])
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			}
		}
	}
	return n1 + n2 - res[n1][n2]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDistance_1(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][0] = i
			dp[0][j] = j
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
		}
	}

	return dp[m][n]
}

func min(nums ...int) int {
	res := 1<<31 - 1

	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}
