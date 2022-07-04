package leetcode31

import (
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	i := n - 1
	// find peak
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}

	if i != 0 {
		j := n - 1
		// i-1: next smaller
		for j > i && nums[j] <= nums[i-1] {
			j--
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
	}
	sort.Ints(nums[i:])
}

func getNextPermutation(nums []int, i int) {
	end := len(nums) - 1
	idx := i
	for j := end; j >= i; j-- {
		if nums[j] > nums[i-1] {
			idx = j
			break
		}
	}
	nums[i-1], nums[idx] = nums[idx], nums[i-1]
	l, r := i, end
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func nextPermutation_1(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}

	sort.Ints(nums)
}
