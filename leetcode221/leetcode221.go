package leetcode221

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}

	maxEdge := 0

	dp := make([][]int, rows)

	for i := range dp {
		dp[i] = make([]int, cols)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxEdge = 1
		}
	}

	for j := 1; j < cols; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxEdge = 1
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(min(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j])

				if dp[i][j] > maxEdge {
					maxEdge = dp[i][j]
				}
			}
		}
	}

	return maxEdge * maxEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
