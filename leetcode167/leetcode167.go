package leetcode167

func twoSum(numbers []int, target int) []int {
	mapper := make(map[int]int, len(numbers))

	max := target - numbers[0]

	for i, v := range numbers {
		if v > max {
			break
		}
		if n := mapper[target-v]; n != 0 {
			return []int{n, i + 1}
		}

		mapper[v] = i + 1
	}
	return nil
}
