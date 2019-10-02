package leetcode16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closet := math.Abs(float64(nums[0] + nums[1] + nums[2] - target))
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				abs := math.Abs(float64(sum - target))
				if abs <= closet {
					closet = abs
					res = sum
				}
			}

		}
	}
	return res
}
