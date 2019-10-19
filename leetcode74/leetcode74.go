package leetcode74

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	row := 0
	for i := m - 1; i >= 0; i-- {
		if matrix[i][0] == target {
			return true
		}
		if matrix[i][0] < target {
			row = i
			break
		}
	}

	l, r, mid := 0, n-1, 0

	for l <= r {
		mid = (l + r) / 2
		switch {
		case matrix[row][mid] < target:
			l = mid + 1
		case matrix[row][mid] > target:
			r = mid - 1
		case matrix[row][mid] == target:
			return true
		}
	}
	return false
}
