package LeetCode_377_CombinationSumIV

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

func combinationSum4_1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
