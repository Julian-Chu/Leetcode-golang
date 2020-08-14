package leetcode685

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1)
	var first, last []int
	for i := range edges {
		f := edges[i][0]
		t := edges[i][1]
		if parents[t] == 0 {
			parents[t] = f
		} else {
			first = []int{parents[t], t}
			last = edges[i]
			edges[i] = nil
			break
		}
	}
	root := parents
	for i := range root {
		root[i] = i
	}

	rootOf := func(t int) int {
		for t != root[t] {
			t = root[t]
		}
		return t
	}

	for _, edge := range edges {
		if edge == nil {
			continue
		}

		f := edge[0]
		t := edge[1]
		r := rootOf(f)

		if r == t {
			if first == nil {
				return edge
			}
			return first
		}

		root[t] = r
	}
	return last

}
