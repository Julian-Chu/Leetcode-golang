package leetcode496

func nextGreaterElement(findNums []int, nums []int) []int {
	if len(findNums) == 0 {
		return []int{}
	}

	m := make(map[int]int)

	for i, num := range nums {
		m[num] = i
	}

	res := make([]int, len(findNums))
	for i, num := range findNums {
		res[i] = -1
		for j := m[num] + 1; j < len(nums); j++ {
			if nums[j] > num {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}
