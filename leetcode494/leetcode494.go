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

	dp := make([]int, 2*sum+1)
	dp[sum] = 1 // origin    0 sum 2*sum+1
	for i := 0; i < len(nums); i++ {
		next := make([]int, 2*sum+1)
		for j := 0; j < len(dp); j++ {
			if dp[j] == 0 {
				continue
			}

			next[j+nums[i]] += dp[j]
			next[j-nums[i]] += dp[j]
		}
		dp = next
	}
	return dp[sum+S]
}
