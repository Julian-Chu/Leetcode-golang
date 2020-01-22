package leetcode200

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	size := len(grid) * len(grid[0])

	x := make([]int, 0, size)
	y := make([]int, 0, size)
	add := func(i, j int) {
		x = append(x, i)
		y = append(y, j)
		grid[i][j] = '0'
	}
	bfs := func(i, j int) int {
		if grid[i][j] == '0' {
			return 0
		}

		add(i, j)

		for len(x) > 0 {
			i := x[0]
			x = x[1:]
			j := y[0]
			y = y[1:]

			if i-1 >= 0 && grid[i-1][j] == '1' {
				add(i-1, j)
			}

			if j-1 >= 0 && grid[i][j-1] == '1' {
				add(i, j-1)
			}

			if i+1 < len(grid) && grid[i+1][j] == '1' {
				add(i+1, j)
			}

			if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
				add(i, j+1)
			}
		}
		return 1

	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res += bfs(i, j)
		}
	}

	return res
}

//func getPos(i int, j int) string {
//	return fmt.Sprintf("%v,%v", i, j)
//}
