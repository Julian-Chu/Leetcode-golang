package leetcode62

func uniquePaths(m int, n int) int {
	var matrix = [100][100]int{}

	for i := 0; i < m; i++ {
		matrix[i][0] = 1
	}

	for j := 0; j < n; j++ {
		matrix[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}

	return matrix[m-1][n-1]

}
