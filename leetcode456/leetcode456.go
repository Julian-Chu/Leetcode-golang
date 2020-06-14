package leetcode456

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	for i := 0; i < len(nums)-2; i++ {
		max, min := -1<<31, 1<<31-1
		if nums[i] < min {
			min = nums[i]
		}

		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j] > max {
				max = nums[j]
			}

			for k := j + 1; k < len(nums); k++ {
				if nums[k] > min && nums[k] < max {
					return true
				}
			}
		}
	}

	return false
}
