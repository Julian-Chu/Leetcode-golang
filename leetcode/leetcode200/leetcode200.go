package leetcode200

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, columns := getMatrixSize(grid)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, columns)
	}

	var islandCount int
	for row := range grid {
		for column := range grid[row] {
			if dfsTriggered(row, column, grid, visited) {
				islandCount++
			}
		}
	}

	return islandCount
}

func getMatrixSize(grid [][]byte) (rows int, columns int) {
	rows = len(grid)
	columns = len(grid[0])
	return
}

func dfsTriggered(row, column int, grid [][]byte, visited [][]bool) bool {
	if visited[row][column] {
		return false
	}

	dfs(row, column, grid, visited)
	if grid[row][column] == '0' {
		return false
	}
	return true
}

func dfs(row, column int, grid [][]byte, visited [][]bool) {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[row]) || visited[row][column] == true {
		return
	}

	if grid[row][column] == '0' {
		visited[row][column] = true
		return
	}

	visited[row][column] = true
	dfs(row, column-1, grid, visited)
	dfs(row, column+1, grid, visited)
	dfs(row-1, column, grid, visited)
	dfs(row+1, column, grid, visited)
	return
}
