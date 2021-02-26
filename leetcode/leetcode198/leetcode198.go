package leetcode198

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := nums[:]
	max := nums[0]
	if n > 1 {
		max = Max(nums[0], nums[1])
	}
	for i := 2; i < n; i++ {
		if i == 2 {
			dp[i] = nums[i] + nums[i-2]
			max = Max(max, dp[i])
			continue
		}
		dp[i] = dp[i-2] + nums[i]
		if dp[i-3] > dp[i-2] {
			dp[i] = dp[i-3] + nums[i]
		}
		max = Max(max, dp[i])
	}

	if n >= 2 && dp[n-2] > dp[n-1] {
		return dp[n-2]
	}
	return dp[n-1]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
