package leetcode35

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)>>1
		switch {
		case nums[mid] == target:
			end = mid
		case nums[mid] < target:
			start = mid
		default:
			end = mid
		}
	}

	if nums[start] >= target {
		return start
	}

	if nums[end] < target {
		return end + 1
	}

	return end
}

//func searchInsert(nums []int, target int) int {
//	n := len(nums)
//	l, r := 0, n-1
//	var mid int
//	for l <= r {
//		mid = (l + r) / 2
//		switch {
//		case nums[mid] < target:
//			l = mid + 1
//		case nums[mid] > target:
//			r = mid - 1
//		default:
//			return mid
//		}
//	}
//
//	if target > nums[mid] {
//		return mid + 1
//	}
//
//	return mid
//}
