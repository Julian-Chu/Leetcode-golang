package leetcode435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := make([][]int, 0, len(intervals))
	end := -1 << 31
	for i := range intervals {
		if intervals[i][0] >= end {
			res = append(res, intervals[i])
			end = intervals[i][1]
		}
	}

	return len(intervals) - len(res)
}
