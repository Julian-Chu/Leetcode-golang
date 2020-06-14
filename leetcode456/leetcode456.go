package leetcode456

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i+1] < nums[i] {
			continue
		}
		max, min := -1<<31, 1<<31-1
		if nums[i] < min {
			min = nums[i]
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
				continue
			}

			if nums[j] > min && nums[j] < max {
				return true
			}
		}
	}

	return false
}
