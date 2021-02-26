package leetcode73

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	zerosInFirstRow := false
	zerosInFirstCol := false

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				if i == 0 {
					zerosInFirstRow = true
				}
				if j == 0 {
					zerosInFirstCol = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := m - 1; i > 0; i-- {
		for j := n - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if zerosInFirstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if zerosInFirstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}
