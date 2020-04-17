package leetcode494

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if (sum+S)%2 == 1 {
		return 0
	}
	target := (sum + S) / 2
	dp := make([]int, 1001)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[target]
}
