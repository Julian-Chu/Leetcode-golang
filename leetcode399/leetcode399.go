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

	var findPath func(start, end string, route map[string]bool) float64
	findPath = func(start, end string, route map[string]bool) float64 {
		nexts, checkStart := edges[start]
		_, checkEnd := edges[end]

		if !(checkStart && checkEnd) {
			return -1
		}

		if start == end {
			return 1
		}
		for k, v := range nexts {
			if k == end {
				return v
			}
			if exist := route[k]; exist {
				continue
			}

			route[k] = true
			res := findPath(k, end, route)
			if res != -1 {
				return v * res
			}
			route[k] = false
		}
		return -1
	}

	res := make([]float64, 0, len(queries))

	for i := range queries {
		start := queries[i][0]
		end := queries[i][1]
		route := make(map[string]bool)
		route[start] = true
		res = append(res, findPath(start, end, route))
	}
	return res
}
