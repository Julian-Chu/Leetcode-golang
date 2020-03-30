package leetcode274

import "sort"

func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	hIdx := 0
	for idx := 1; idx <= n; idx++ {
		gt := 0
		eq := 0
		for i := range citations {
			switch {
			case citations[i] > idx:
				gt++
			case citations[i] == idx:
				eq++
			}
		}
		if gt+eq >= idx {
			hIdx = idx
		}

	}
	return hIdx
}
