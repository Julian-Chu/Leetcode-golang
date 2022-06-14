package leetcode63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m := len(obstacleGrid) - 1
	n := len(obstacleGrid[0]) - 1
	res := make([][]int, 100)
	for i := range res {
		res[i] = make([]int, 100)
	}

	for i := 0; i < m+1; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		res[i][0] = 1
	}

	for j := 0; j < n+1; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		res[0][j] = 1
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if obstacleGrid[i][j] == 1 {
				res[i][j] = 0
				continue
			}
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}

	return res[m][n]
}

func uniquePathsWithObstacles_1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n && grid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 0; i < m && grid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if grid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}

	return dp[m-1][n-1]
}
