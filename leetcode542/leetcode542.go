package leetcode542

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	MIN := m * n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}

			matrix[i][j] = MIN
			if 0 <= i-1 {
				matrix[i][j] = min(matrix[i][j], matrix[i-1][j]+1)
			}

			if 0 <= j-1 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j-1]+1)
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}

			if i+1 < m {
				matrix[i][j] = min(matrix[i][j], matrix[i+1][j]+1)
			}

			if j+1 < n {
				matrix[i][j] = min(matrix[i][j], matrix[i][j+1]+1)
			}
		}
	}
	return matrix
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}
