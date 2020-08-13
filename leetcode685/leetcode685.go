package leetcode685

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	var first, last []int

	for k := range edges {
		p, c := edges[k][0], edges[k][1]
		if parent[c] == 0 {
			parent[c] = p
		} else {
			first = []int{parent[c], c}
			last = edges[k]
			edges[k] = nil
			break
		}
	}

	root := parent

	for i := 0; i <= n; i++ {
		root[i] = i
	}

	rootOf := func(i int) int {
		for i != root[i] {
			i = root[i]
		}
		return i
	}

	for _, edge := range edges {
		if edge == nil {
			continue
		}

		p := edge[0]
		c := edge[1]
		r := rootOf(p)
		if r == c {
			if first == nil {
				return edge
			}
			return first
		}
		root[c] = r
	}

	return last
}
