package leetcode81

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	offset := 0
	for i := range nums {
		if i > 0 && nums[i] < nums[i-1] {
			offset = i
		}
	}

	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		newMid := transform(mid, n, offset)
		switch {
		case target > nums[newMid]:
			l = mid + 1
		case target < nums[newMid]:
			r = mid - 1
		default:
			return true
		}
	}
	return false
}

func transform(pos, len, offset int) int {
	return (pos + offset) % len
}
