package leetcode435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	end := -1 << 31
	for i := range intervals {
		if intervals[i][0] >= end {
			end = intervals[i][1]
		} else {
			res++
		}
	}

	return res
}
