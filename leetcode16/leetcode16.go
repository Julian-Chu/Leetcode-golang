package leetcode16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	diff := math.Abs(float64(math.MaxInt32))
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			delta := sum - target
			abs := math.Abs(float64(delta))

			if abs < diff {
				diff = abs
				res = sum
			}

			if delta < 0 {
				l++
			} else {
				r--
			}

		}
	}
	return res
}
