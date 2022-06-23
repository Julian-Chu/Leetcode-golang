package LeetCode_724_FindPivotIndex

func pivotIndex(nums []int) int {
	n := len(nums)

	leftSum := 0
	rightSum := 0
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for i := 0; i < n; i++ {
		leftSum += nums[i]
		rightSum = sum - leftSum + nums[i]
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}
