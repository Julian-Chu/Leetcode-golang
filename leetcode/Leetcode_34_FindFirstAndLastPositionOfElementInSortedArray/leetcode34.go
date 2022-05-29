package Leetcode_34_FindFirstAndLastPositionOfElementInSortedArray

//func searchRange(nums []int, target int) []int {
//	n := len(nums)
//	if n == 0 {
//		return []int{-1, -1}
//	}
//	l, r := 0, n-1
//	mid := 0
//	for l < r {
//		mid = (l + r) / 2
//		switch {
//		case target > nums[mid]:
//			l = mid + 1
//		case target < nums[mid]:
//			r = mid - 1
//		default:
//			l, r = mid, mid
//		}
//	}
//	mid = (l + r) / 2
//	if nums[mid] != target {
//		return []int{-1, -1}
//	}
//
//	for l = mid; l >= 0; l-- {
//		if nums[l] != target {
//			break
//		}
//	}
//	l++
//	for r = mid; r < n; r++ {
//		if nums[r] != target {
//			break
//		}
//	}
//	r--
//
//	return []int{l, r}
//}

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
