package leetcode456

func find132pattern(nums []int) bool {

	numk := -1 << 31

	stack := make([]int, 0, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < numk {
			return true
		}

		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			numk = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums[i])
	}
	return false
}
