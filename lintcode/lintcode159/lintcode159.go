package lintcode159

/**
 * @param nums: a rotated sorted array
 * @return: the minimum number in the array
 */
func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1 << 31
	}

	lastNum := nums[len(nums)-1]
	start, end := 0, len(nums)

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > lastNum {
			start = mid
		} else {
			end = mid
		}
	}

	return min(nums[start], nums[end])

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
