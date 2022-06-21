package LeetCode_503_NextGreaterElementII

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

func nextGreaterElements_1(nums []int) []int {
	var stack []int
	n := len(nums)

	res := make([]int, n)
	for i := range nums {
		res[i] = -1
	}
	for i := 0; i < n*2; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = nums[i%n]
		}

		stack = append(stack, i%n)
	}

	return res
}
