package leetcode74

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, m-1

	var mid int
	for l <= r {
		mid = (l + r) / 2
		switch {
		case matrix[mid][0] < target:
			l = mid + 1
		case matrix[mid][0] > target:
			r = mid - 1
		case matrix[mid][0] == target:
			return true
		}
	}

	row := mid
	if matrix[mid][0] > target && mid > 0 {
		row = mid - 1
	}

	l, r, mid = 0, n-1, 0

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
