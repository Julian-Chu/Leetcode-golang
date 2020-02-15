package leetcode283

func moveZeroes(nums []int) {
	l := len(nums)
	if l == 0 || l == 1 {
		return
	}

	for i := l - 2; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		if nums[i+1] == 0 {
			continue
		}

		for j := i; j < l-1; j++ {
			for nums[j+1] != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return
}
