package binarysearch

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start+1 < end {
		mid := start + (end-start)/2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			start = mid
		case nums[mid] > target:
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	return -1
}

func bisect_first_gte(sub []int, num int) int {
	if len(sub) == 0 {
		return 0
	}

	start, end := 0, len(sub)-1

	for start+1 < end {
		mid := start + (end-start)>>1
		if sub[mid] >= num {
			end = mid
		} else {
			start = mid
		}
	}

	if sub[start] >= num {
		return start
	}
	return end
}
