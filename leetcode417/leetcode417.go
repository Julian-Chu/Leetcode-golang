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

	var pQueue = [][]int{}
	var aQueue = [][]int{}

	for i := 0; i < m; i++ {
		p[i][0] = true
		pQueue = append(pQueue, []int{i, 0})
		a[i][n-1] = true
		aQueue = append(aQueue, []int{i, n - 1})
	}

	for j := 0; j < n; j++ {
		p[0][j] = true
		pQueue = append(pQueue, []int{0, j})
		a[m-1][j] = true
		aQueue = append(aQueue, []int{m - 1, j})

	}

	steps := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var res [][]int

	bfs := func(queue [][]int, rec [][]bool) {
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			for _, step := range steps {
				i, j := pos[0]+step[0], pos[1]+step[1]
				if i >= 0 && i < m && j >= 0 && j < n && !rec[i][j] && matrix[pos[0]][pos[1]] <= matrix[i][j] {
					rec[i][j] = true
					queue = append(queue, []int{i, j})
				}
			}
		}
	}

	bfs(pQueue, p)
	bfs(aQueue, a)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
