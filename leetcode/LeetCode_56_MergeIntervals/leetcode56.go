package merge_intervals

import "sort"

func merge_1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][1]

	res := [][]int{}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end = max(intervals[i][1], end)
			continue
		}

		res = append(res, []int{start, end})
		start = intervals[i][0]
		end = intervals[i][1]
	}

	res = append(res, []int{start, end})
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

// {2, 3},  {8, 9}, {1, 10}
// {1, 10}
func merge(intervals []Interval) []Interval {
	res := Sort(intervals)
	last := len(res) - 1
	for i := last; i > 0; i-- {
		if res[i-1].End < res[i].Start || res[i-1].Start > res[i].End {
			continue
		}

		var newStart, newEnd int
		newStart = res[i].Start
		if res[i-1].Start <= res[i].Start {
			newStart = res[i-1].Start
		}
		newEnd = res[i].End
		if res[i-1].End >= res[i].End {
			newEnd = res[i-1].End
		}
		res[i-1] = Interval{newStart, newEnd}

		res = append(res[:i], res[(i+1):]...)

	}
	i := 1
	for i < len(res) {
		if res[i-1].End < res[i].Start || res[i-1].Start > res[i].End {
			i++
			continue
		}

		var newStart, newEnd int
		newStart = res[i].Start
		if res[i-1].Start <= res[i].Start {
			newStart = res[i-1].Start
		}
		newEnd = res[i].End
		if res[i-1].End >= res[i].End {
			newEnd = res[i-1].End
		}
		res[i-1] = Interval{newStart, newEnd}

		res = append(res[:i], res[(i+1):]...)
	}
	return res
}

func Sort(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	start := make([]Interval, 0)
	end := make([]Interval, 0)
	m := intervals[int(len(intervals)/2)]

	for _, item := range intervals {
		switch {
		case item.Start < m.Start:
			start = append(start, item)
		case item.Start == m.Start:
			if item.End < m.End {
				start = append(start, item)
			} else if item.End > m.End {
				end = append(end, item)
			}
		case item.Start > m.Start:
			end = append(end, item)
		}
	}

	start, end = Sort(start), Sort(end)
	start = append(start, m)
	start = append(start, end...)
	return start

}
