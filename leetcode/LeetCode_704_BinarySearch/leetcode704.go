package LeetCode_704_BinarySearch

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

//func search(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//
//	for left <= right {
//		mid := left + (right-left)>>1
//		switch {
//		case nums[mid] > target:
//			right = mid - 1
//		case nums[mid] < target:
//			left = mid + 1
//		default:
//			return mid
//		}
//	}
//
//	return -1
//}
