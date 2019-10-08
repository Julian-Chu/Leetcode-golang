package leetcode35

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	if target <= nums[l] {
		return 0
	}

	if target > nums[r] {
		return n
	}

	var mid int
loop:
	for l <= r {
		mid = (l + r) / 2
		switch {
		case nums[mid] < target:
			if target < nums[mid+1] {
				return mid + 1
			}
			l = mid + 1
		case nums[mid] > target:
			if nums[mid-1] < target {
				return mid
			}
			r = mid - 1
		default:
			break loop
		}
	}
	return mid
}
