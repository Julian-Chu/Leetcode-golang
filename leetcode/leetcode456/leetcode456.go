package leetcode456

func find132pattern(nums []int) bool {
	if nums == nil || len(nums) < 3 {
		return false
	}

	stack := []int{nums[len(nums)-1]}

	aux := make([]int, len(nums))
	aux[0] = nums[0]

	for i := 0; i < len(nums); i++ {
		aux[i] = min(nums[i], aux[i-1])
	}

	for j := len(nums) - 2; j >= 0; j-- {
		if nums[j] > aux[j] {
			for len(stack) != 0 && stack[len(stack)-1] <= aux[j] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) != 0 && stack[len(stack)-1] < nums[j] {
				return true
			}

			stack = append(stack, nums[j])
		}
	}

	return false
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
