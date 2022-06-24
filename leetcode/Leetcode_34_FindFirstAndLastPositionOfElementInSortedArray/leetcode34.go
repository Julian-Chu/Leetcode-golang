package Leetcode_34_FindFirstAndLastPositionOfElementInSortedArray

func searchRange_1(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	start, end := 0, n-1

	for start+1 < end {
		mid := start + (end-start)>>1
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] != target && nums[end] != target {
		return []int{-1, -1}
	}

	if nums[start] != target {
		start = end
	}

	for i := start; i < n; i++ {
		if nums[i] != target {
			break
		}
		end = i
	}

	return []int{start, end}
}

func searchRange(nums []int, target int) []int {
	return []int{
		search(nums, target, true),
		search(nums, target, false),
	}
}
func search(nums []int, target int, isFirst bool) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start+1 < end {
		mid := start + (end-start)/2
		switch {
		case nums[mid] == target:
			if isFirst {
				end = mid
			} else {
				start = mid
			}
		case nums[mid] < target:
			start = mid
		default:
			end = mid
		}
	}

	if isFirst {
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}

	} else {
		if nums[end] == target {
			return end
		}

		if nums[start] == target {
			return start
		}
	}

	return -1
}
