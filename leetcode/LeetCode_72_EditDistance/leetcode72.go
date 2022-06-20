package LeetCode_72_EditDistance

func minDistance(word1 string, word2 string) int {
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

			dp[i][j] = min(
				dp[i-1][j],
				dp[i][j-1],
				dp[i-1][j-1],
			) + 1
		}
	}

	return dp[m][n]
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}
