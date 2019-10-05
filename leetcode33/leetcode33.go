package leetcode33

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	offset := l

	l, r = 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		midOffset := (mid + offset) % len(nums)
		if target > nums[midOffset] {
			l = mid + 1
		} else if target < nums[midOffset] {
			r = mid - 1
		} else {
			return midOffset
		}
	}

	return -1
}
