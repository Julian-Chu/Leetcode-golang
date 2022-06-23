package LeetCode_283_MoveZeros

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

func moveZeroes_1(nums []int) {
	slow := 0

	for i := range nums {
		if nums[i] == 0 {
			continue
		}

		nums[slow], nums[i] = nums[i], nums[slow]
		slow++
	}
}
