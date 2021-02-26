package leetcode503

func nextGreaterElements(nums []int) []int {
	var stack Stack

	for i := len(nums) - 2; i >= 0; i-- {
		for len(stack) > 0 && stack.top() <= nums[i] {
			stack.pop()
		}
		stack.push(nums[i])
	}

	greater := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack.top() <= nums[i] {
			stack.pop()
		}

		if len(stack) == 0 {
			greater[i] = -1
		} else {
			greater[i] = stack.top()
		}
		stack.push(nums[i])
	}
	return greater
}

type Stack []int

func (s *Stack) push(val int) {
	*s = append(*s, val)
}

func (s *Stack) pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack) top() int {
	return (*s)[len(*s)-1]
}
