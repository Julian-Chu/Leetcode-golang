package leetcode417

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	m, n := len(matrix), len(matrix[0])

	p := make([][]bool, m)
	a := make([][]bool, m)

	for i := range p {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}
	var dfs func(i, j int, visited [][]bool)

	steps := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	dfs = func(i, j int, visited [][]bool) {
		visited[i][j] = true

		for _, step := range steps {
			a, b := i+step[0], j+step[1]
			if a < 0 || b < 0 || a >= m || b >= n {
				continue
			}

			if matrix[i][j] <= matrix[a][b] && visited[a][b] == false {
				dfs(a, b, visited)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, p)
		dfs(i, n-1, a)
	}

	for j := 0; j < n; j++ {
		dfs(0, j, p)
		dfs(m-1, j, a)
	}

	var res = [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] == true && a[i][j] == true {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
