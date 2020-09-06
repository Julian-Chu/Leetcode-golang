package leetcode215

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		newPivot := partition(l, r, nums)
		if newPivot == k-1 {
			return nums[newPivot]
		} else if newPivot > k-1 {
			r = newPivot - 1
		} else {
			l = newPivot + 1
		}
	}
	return -1
}

func partition(l int, r int, nums []int) int {
	pv := nums[r]
	newPivot := l
	for i := l; i < r; i++ {
		if nums[i] > pv {
			nums[i], nums[newPivot] = nums[newPivot], nums[i]
			newPivot++
		}
	}

	nums[r], nums[newPivot] = nums[newPivot], nums[r]
	return newPivot
}
