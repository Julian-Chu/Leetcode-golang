package leetcode436

import "sort"

func findRightInterval(intervals [][]int) []int {
	if len(intervals) == 0 {
		return []int{}
	}
	if len(intervals) == 1 {
		return []int{-1}
	}

	m := make(map[int]int, len(intervals))

	for i := range intervals {
		m[intervals[i][0]] = i
	}

	starts := make([]int, 0, len(intervals))
	for key, _ := range m {
		starts = append(starts, key)
	}
	sort.Ints(starts)

	findRight := func(interval []int) (index int) {
		target := interval[1]
		if target > starts[len(starts)-1] {
			return -1
		}

		l, r := 0, len(starts)-1

		for l < r {
			mid := l + (r-l)>>1
			if starts[mid] > target {
				r = mid
			} else if starts[mid] < target {
				l = mid + 1
			} else {
				l = mid
				r = mid
			}
		}
		return m[starts[r]]
	}

	res := make([]int, len(intervals))

	for i := range intervals {
		res[i] = findRight(intervals[i])
	}

	return res
}
