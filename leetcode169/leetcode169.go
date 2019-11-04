package leetcode169

func majorityElement(nums []int) int {
	m := make(map[int]int)

	moreThan := len(nums) / 2
	for _, v := range nums {
		m[v]++
		if m[v] > moreThan {
			return v
		}
	}
	return 0
}
