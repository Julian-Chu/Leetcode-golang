package leetcode452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	start := points[0][0]
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if end >= points[i][0] {
			if start < points[i][0] {
				start = points[i][0]
			}

			if end > points[i][1] {
				end = points[i][1]
			}
		} else {
			res++
			start = points[i][0]
			end = points[i][1]
		}

	}

	return res
}
