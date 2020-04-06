package leetcode463

func islandPerimeter(grid [][]int) int {
	l := len(grid)
	if l == 0 {
		return 0
	}
	w := len(grid[0])
	if w == 0 {
		return 0
	}
	perimeter := 0
	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 || grid[i-1][j] == 0 {
				perimeter++
			}

			if i == l-1 || grid[i+1][j] == 0 {
				perimeter++
			}

			if j == 0 || grid[i][j-1] == 0 {
				perimeter++
			}

			if j == w-1 || grid[i][j+1] == 0 {
				perimeter++
			}
		}
	}

	return perimeter
}
