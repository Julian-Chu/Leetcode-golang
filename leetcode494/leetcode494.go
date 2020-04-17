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

	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		next := make(map[int]int)
		for sum, cnt := range dp {
			next[sum+nums[i]] += cnt
			next[sum-nums[i]] += cnt
		}
		dp = next
	}
	return dp[S]
}
