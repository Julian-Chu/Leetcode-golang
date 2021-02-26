package leetcode399

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	edges := make(map[string]map[string]float64)

	for i := range equations {
		start := equations[i][0]
		end := equations[i][1]
		val := values[i]
		if _, ok := edges[start]; !ok {
			edges[start] = make(map[string]float64)
		}
		if _, ok := edges[end]; !ok {
			edges[end] = make(map[string]float64)
		}

		edges[start][end] = val
		edges[end][start] = 1 / val
	}

	bfs := func(graph map[string]map[string]float64, a, b string) float64 {
		if _, ok := graph[a]; !ok {
			return -1.0
		}

		if _, ok := graph[b]; !ok {
			return -1.0
		}

		if a == b {
			return 1.0
		}

		type entry struct {
			s string
			f float64
		}
		isVisited := make(map[string]bool)
		queue := []entry{{a, 1}}
		for len(queue) > 0 {
			e := queue[0]
			queue = queue[1:]
			if e.s == b {
				return e.f
			}
			if isVisited[e.s] {
				continue
			}

			isVisited[e.s] = true

			for k, v := range graph[e.s] {
				queue = append(queue, entry{k, v * e.f})
			}
		}
		return -1
	}

	res := make([]float64, len(queries))

	for i := range queries {
		start := queries[i][0]
		end := queries[i][1]
		res[i] = bfs(edges, start, end)
	}
	return res
}
