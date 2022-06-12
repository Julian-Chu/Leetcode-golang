package leetcode53

func maxSubArray(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	max := nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray_greedy(nums []int) int {
	res := -1 << 31
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum > res {
			res = sum
		}

		if sum <= 0 {
			sum = 0
		}
	}
	return res
}

func maxSubArray_dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
