package leetcode310

func findMinHeightTrees(n int, edges [][]int) []int {
	nodes := make([][]int, n)

	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}

	heights := make([]int, n)
	for i := range nodes {
		m := make([]bool, n)
		m[i] = true
		queue := []int{i}
		for len(queue) > 0 {
			heights[i]++
			nextq := make([]int, 0, len(nodes[i]))
			for _, node := range queue {
				for _, next := range nodes[node] {
					if m[next] == false {
						nextq = append(nextq, next)
						m[next] = true
					}
				}
			}
			queue = nextq
		}
	}

	min := heights[0]
	for _, h := range heights {
		if h < min {
			min = h
		}
	}

	res := make([]int, 0)

	for i, h := range heights {
		if h == min {
			res = append(res, i)
		}
	}
	return res
}
