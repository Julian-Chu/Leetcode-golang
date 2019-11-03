package leetcode167

func twoSum(numbers []int, target int) []int {
	mapper := make(map[int]int)

	max := target - numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			break
		}
		if _, ok := mapper[numbers[i]]; !ok {
			mapper[numbers[i]] = i
		}
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > target {
			break
		}

		nextVal := target - numbers[i]
		if nextVal == numbers[i] {
			return []int{i + 1, i + 2}
		}

		if idx, ok := mapper[nextVal]; ok {
			return []int{i + 1, idx + 1}
		}
	}
	return []int{0, 0}
}
