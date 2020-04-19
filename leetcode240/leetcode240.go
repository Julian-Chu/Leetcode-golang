package leetcode240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])

	i, j := 0, n-1

	for i < m && j >= 0 {
		val := matrix[i][j]

		switch {
		case val > target:
			j--
		case val < target:
			i++
		default:
			return true
		}
	}

	return false
}
