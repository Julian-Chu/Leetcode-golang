package leetcode496

func nextGreaterElement(findNums []int, nums []int) []int {
	if len(findNums) == 0 {
		return []int{}
	}

	m := make(map[int]int, len(nums))
	greaterMap := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
		greaterMap[num] = -1
	}

	stack := make([]int, 0, len(nums))
	stack = append(stack, nums[0])
	for i := 1; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			greaterMap[n] = nums[i]
		}
		stack = append(stack, nums[i])
	}

	res := make([]int, len(findNums))

	for i := range findNums {
		res[i] = greaterMap[findNums[i]]
	}

	return res
}
