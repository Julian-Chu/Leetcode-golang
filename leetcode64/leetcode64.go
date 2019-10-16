package leetcode64

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	matrix[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		matrix[i][0] = matrix[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		matrix[0][j] = matrix[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			previousMin := matrix[i-1][j]
			if matrix[i-1][j] > matrix[i][j-1] {
				previousMin = matrix[i][j-1]
			}
			matrix[i][j] = grid[i][j] + previousMin
		}
	}
	return matrix[m-1][n-1]
}
