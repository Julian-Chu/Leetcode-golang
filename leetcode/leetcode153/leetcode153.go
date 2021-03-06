package leetcode153

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			min = nums[i]
		}
	}

	return min
}
