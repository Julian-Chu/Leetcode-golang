package LeetCode_496_NextGreaterElementI

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	m := map[int]int{}
	for i, num := range nums2 {
		for len(stack) > 0 && num > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			m[nums2[top]] = num
		}

		stack = append(stack, i)
	}

	res := make([]int, len(nums1))

	for i, num := range nums1 {
		res[i] = -1
		if v, ok := m[num]; ok {
			res[i] = v
		}
	}

	return res
}

func nextGreaterElement_1(nums1 []int, nums2 []int) []int {
	stack := []int{}
	m := map[int]int{}
	res := make([]int, len(nums1))

	for i, num := range nums1 {
		res[i] = -1
		m[num] = i
	}

	for i, num := range nums2 {
		for len(stack) > 0 && num > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if idx, ok := m[nums2[top]]; ok {
				res[idx] = num
			}
		}

		stack = append(stack, i)
	}

	return res
}
