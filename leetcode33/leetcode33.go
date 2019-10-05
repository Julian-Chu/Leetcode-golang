package leetcode33

func search(nums []int, target int) int {
	if len(nums) <= 2 {
		for i, v := range nums {
			if v == target {
				return i
			}
		}
		return -1
	}
	l, r := 0, len(nums)-1
	mid := (l + r) / 2

	for mid > l && nums[mid] < nums[l] {
		mid--
	}

	for mid < r && nums[mid+1] > nums[r] {
		mid++
	}

	if mid == l {
		if target > nums[l] || target < nums[l+1] {
			return -1
		}
	}

	if target < nums[l] {
		l = mid + 1
	} else if target > nums[r] {
		r = mid
	}

	mid = (l + r) / 2
	for l < r {
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
		mid = (l + r) / 2
	}

	if nums[mid] != target {
		return -1
	}
	return mid
}
