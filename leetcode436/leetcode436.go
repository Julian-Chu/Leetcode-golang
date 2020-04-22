package leetcode436

import "sort"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{-1}
	}

	m := make(map[int]int, n)

	starts := make([]int, n)
	for i := range intervals {
		m[intervals[i][0]] = i
		starts[i] = intervals[i][0]
	}

	sort.Ints(starts)

	res := make([]int, n)

	for i, v := range intervals {
		idx := sort.SearchInts(starts, v[1])
		if idx < n {
			res[i] = m[starts[idx]]
		} else {
			res[i] = -1
		}
	}

	return res
}
