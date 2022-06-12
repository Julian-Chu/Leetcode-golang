package LeetCode_1005_MaximizeSumOfArrayAfterKNegations

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
		sum += nums[i]
	}
	k %= 2
	if k == 1 {
		min := 1<<31 - 1
		for i := range nums {
			if nums[i] < min {
				min = nums[i]
			}
		}

		sum -= 2 * min
	}
	return sum
}

func largestSumAfterKNegations_sortByAbs(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	sum := 0
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
		sum += nums[i]
	}
	k %= 2
	if k == 1 {
		sum -= 2 * nums[len(nums)-1]
	}
	return sum
}
