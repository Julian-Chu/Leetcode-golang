package leetcode368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	n := len(nums)
	dp := make([][]int, n)

	sort.Ints(nums)
	for idx := range nums {
		if dp[idx] == nil {
			dp[idx] = []int{nums[idx]}
		}
		for i := 0; i < idx; i++ {
			if nums[i]<<1 > nums[idx] {
				break
			}
			if nums[idx]%nums[i] == 0 && len(dp[idx]) < len(dp[i])+1 {
				dp[i] = dp[i][:len(dp[i]):len(dp[i])]
				dp[idx] = append(dp[i], nums[idx])
			}
		}
	}

	maxIdx := 0
	for i := range dp {
		if len(dp[maxIdx]) < len(dp[i]) {
			maxIdx = i
		}
	}
	return dp[maxIdx]
}
