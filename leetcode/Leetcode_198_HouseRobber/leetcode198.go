package Leetcode_198_HouseRobber

func rob(nums []int) int {
	n := len(nums)

	switch {
	case n < 1:
		return 0
	case n == 1:
		return nums[0]
	case n == 2:
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	case n == 3:
		sum := nums[0] + nums[2]
		if sum > nums[1] {
			return sum
		}
		return nums[1]

	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[2] + nums[0]

	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
	}

	return max(dp[n-1], dp[n-2])
}

func rob_1(nums []int) int {
	n := len(nums)

	switch {
	case n < 1:
		return 0
	case n == 1:
		return nums[0]
	case n == 2:
		return max(nums[0], nums[1])
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i]) // max(not rob i, rob i)
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
