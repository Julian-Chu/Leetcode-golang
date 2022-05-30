package LeetCode_209_MinimumSizeSubarraySum

import "math"

func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	sum := 0
	idx := 0
	subLen := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			subLen = i - idx + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[idx]
			idx++
		}
	}
	if result != math.MaxInt {
		return result
	}
	return 0
}
