package sort

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}
	res := make([]int, 0, len(nums))
	for key := range m {
		res = append(res, key)
	}
	sort.Slice(res, func(i, j int) bool {
		return m[res[i]] > m[res[j]]
	})
	return res[:k]
}
