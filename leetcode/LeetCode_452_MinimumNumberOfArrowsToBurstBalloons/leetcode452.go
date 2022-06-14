package LeetCode_452_MinimumNumberOfArrowsToBurstBalloons

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if end >= points[i][0] {
			continue
		} else {
			res++
			end = points[i][1]
		}

	}

	return res
}

func findMinArrowShots_1(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 1
	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			res++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	return res
}

func min(a, b int) int {
	if a >= b {
		return b
	}

	return a
}