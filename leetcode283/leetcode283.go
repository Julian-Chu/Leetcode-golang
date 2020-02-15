package leetcode283

func moveZeroes(nums []int) {
	replace := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[replace] = nums[i]
			replace++
		}
	}

	for i := replace; i < len(nums); i++ {
		nums[i] = 0
	}
}
