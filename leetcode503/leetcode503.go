package leetcode503

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res[i] = -1
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = nums[i]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = nums[i]
		}
	}

	return res
}
