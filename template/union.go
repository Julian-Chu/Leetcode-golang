package template

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)

	parent := make([]int, n+1)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var i int
	var e []int
	for i, e = range edges {
		f, t := e[0], e[1]
		pf := find(parent, f)
		pt := find(parent, t)
		if pf == pt { // last edge will make connected
			break
		}
		parent[pf] = pt
	}
	return edges[i]
}

func find(parent []int, f int) int {
	for f != parent[f] {
		f = parent[f]
	}
	return f
}
