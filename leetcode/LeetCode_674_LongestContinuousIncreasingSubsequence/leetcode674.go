package LeetCode_674_LongestContinuousIncreasingSubsequence

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}

func findLengthOfLCIS_greedy(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res, count := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > res {
			res = count
		}
	}
	return res
}
