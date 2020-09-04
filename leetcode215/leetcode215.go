package leetcode215

import "sort"

func findKthLargest(nums []int, k int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	keys := make([]int, 0, len(cnt))

	for key := range cnt {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		k -= cnt[keys[i]]
		if k <= 0 {
			return keys[i]
		}
	}

	return -1
}
