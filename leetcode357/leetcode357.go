package leetcode357

func countNumbersWithUniqueDigits(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10

	nums := []int{9, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	if n > 10 {
		n = 10
	}
	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			dp[i] *= nums[j]
		}
		dp[i] += dp[i-1]
	}

	return dp[n]
}
