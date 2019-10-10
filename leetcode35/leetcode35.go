package leetcode35

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		switch {
		case nums[mid] < target:
			l = mid + 1
		case nums[mid] > target:
			r = mid - 1
		default:
			return mid
		}
	}

	if target > nums[mid] {
		return mid + 1
	}

	return mid
}
