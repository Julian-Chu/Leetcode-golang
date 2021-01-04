package lintcode168

/**
 * @param nums: A list of integer
 * @return: An integer, maximum coins
 */
func maxCoins(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}

	nums = append(nums, []int{1, 1}...)
	copy(nums[1:], nums)
	nums[0] = 1

	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for l := 3; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
