package leetcode973

func kClosest(points [][]int, K int) [][]int {
	lo, hi := 0, len(points)-1

	for lo < hi {
		pivot := partition(points, lo, hi)
		if pivot > K-1 {
			hi = pivot - 1
		} else if pivot < K-1 {
			lo = pivot + 1
		} else {
			break
		}
	}
	return points[:K]
}

func partition(points [][]int, lo int, hi int) int {
	newPivot := lo
	val := dis(points[hi])
	for i := lo; i <= hi-1; i++ {
		if dis(points[i]) <= val {
			points[i], points[newPivot] = points[newPivot], points[i]
			newPivot++
		}
	}
	points[newPivot], points[hi] = points[hi], points[newPivot]
	return newPivot
}

func dis(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
