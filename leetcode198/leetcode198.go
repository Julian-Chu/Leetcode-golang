package leetcode198

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i < 2 {
			dp[i] = nums[i]
			continue
		}
		if i == 2 {
			dp[i] = nums[i] + nums[i-2]
			continue
		}
		dp[i] = dp[i-2] + nums[i]
		if dp[i-3] > dp[i-2] {
			dp[i] = dp[i-3] + nums[i]
		}
	}

	if n >= 2 && dp[n-2] > dp[n-1] {
		return dp[n-2]
	}
	return dp[n-1]
}
