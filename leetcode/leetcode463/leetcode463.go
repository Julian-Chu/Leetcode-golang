package leetcode463

func islandPerimeter(grid [][]int) int {
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}
	m, n := len(grid), len(grid[0])

	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4

			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					res--
				}
			}
		}
	}

	return res
}