package LintCode458

func lastPosition(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start+1 < end {

		mid := start + (end-start)/2
		switch {
		case nums[mid] == target:
			start = mid
		case nums[mid] < target:
			start = mid
		default:
			end = mid
		}
	}

	if nums[end] == target {
		return end
	}

	if nums[start] == target {
		return start
	}

	return -1
}
