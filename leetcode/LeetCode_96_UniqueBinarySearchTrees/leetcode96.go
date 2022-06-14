package LeetCode_96_UniqueBinarySearchTrees

func numTrees(n int) int {
	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
