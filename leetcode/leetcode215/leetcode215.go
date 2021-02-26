package leetcode215

import "sort"

func findKthLargest(nums []int, k int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		newPivot := partition(nums, lo, hi)
		switch {
		case newPivot == k-1:
			return nums[k-1]
		case newPivot > k-1:
			hi = newPivot - 1
		default:
			lo = newPivot + 1
		}
	}
	return -1
}

func partition(nums []int, lo int, hi int) int {
	pv := nums[hi]
	newPivot := lo
	for i := lo; i < hi; i++ {
		if nums[i] > pv {
			nums[i], nums[newPivot] = nums[newPivot], nums[i]
			newPivot++
		}
	}
	nums[hi], nums[newPivot] = nums[newPivot], nums[hi]
	return newPivot
}

func findKthLargest_1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
