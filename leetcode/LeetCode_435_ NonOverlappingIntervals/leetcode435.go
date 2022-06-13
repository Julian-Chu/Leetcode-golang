package LeetCode_435__NonOverlappingIntervals

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

func eraseOverlapIntervals_1(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}

	return len(intervals) - count
}
