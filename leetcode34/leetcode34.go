package leetcode34

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	l, r := 0, n-1
	mid := 0
	for l < r {
		mid = (l + r) / 2
		switch {
		case target > nums[mid]:
			l = mid + 1
		case target < nums[mid]:
			r = mid - 1
		default:
			l, r = mid, mid
		}
	}
	mid = (l + r) / 2
	if nums[mid] != target {
		return []int{-1, -1}
	}

	for l = mid; l >= 0; l-- {
		if nums[l] != target {
			break
		}
	}
	l++
	for r = mid; r < n; r++ {
		if nums[r] != target {
			break
		}
	}
	r--

	return []int{l, r}
}
