package leetcode1584

import "math"

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dist := make([]int, n)
	curr, res := 0, 0
	set := make(map[int]bool)

	for i := range dist {
		dist[i] = math.MaxInt64
	}

	for i := 0; i < n-1; i++ {
		set[curr] = true
		dist[curr] = math.MaxInt64

		x, y := points[curr][0], points[curr][1]

		for j := 0; j < n; j++ {
			if _, ok := set[j]; ok {
				continue
			}
			x1, y1 := points[j][0], points[j][1]
			d := abs(x-x1) + abs(y-y1)
			if dist[j] > d {
				dist[j] = d
			}
		}

		minDist, nearestPoint := math.MaxInt64, -1
		for i, d := range dist {
			if d < minDist {
				minDist = d
				nearestPoint = i
			}
		}

		res += minDist
		curr = nearestPoint
	}
	return res
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
