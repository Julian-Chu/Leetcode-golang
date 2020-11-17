package leetcode1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		key := target - num
		if idx, ok := m[key]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{-1, -1}
}
