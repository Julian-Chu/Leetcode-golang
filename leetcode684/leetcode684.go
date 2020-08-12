package leetcode684

func findRedundantConnection(edges [][]int) []int {
	graph := make(map[int][]int)

	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]int, 0, len(edges)-1)
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])

		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make([]int, 0, len(edges)-1)
		}
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	anySingleNode := true

	for anySingleNode {
		for k := range graph {
			if len(graph[k]) != 1 {
				continue
			}
			v := graph[k][0]
			delete(graph, k)
			for i := range graph[v] {
				if graph[v][i] == k {
					graph[v][i], graph[v][len(graph[v])-1] = graph[v][len(graph[v])-1], graph[v][i]
					graph[v] = graph[v][:len(graph[v])-1]
					break
				}
			}
		}

		anySingleNode = false
		for k := range graph {
			if len(graph[k]) == 1 {
				anySingleNode = true
				break
			}
		}
	}

	loopMap := make(map[int]bool)

	for k, v := range graph {
		if len(v) != 1 {
			loopMap[k] = true
		}
	}

	for i := len(edges) - 1; i >= 0; i-- {
		if loopMap[edges[i][0]] && loopMap[edges[i][1]] {
			return edges[i]
		}
	}

	return nil
}
