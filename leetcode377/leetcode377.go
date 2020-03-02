package leetcode377

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := range nums {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}
