package template

var (
	n      = 1001 // number of nodes between 3 and 1000
	parent = make([]int, 1001)
)

// init
func initialize() {
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}

// find the root
func find_1(u int) int {
	//if u == parent[u] {
	//	return u
	//}
	//parent[u] = find_1(parent[u])
	//return parent[u]

	root := parent[u]
	for root != parent[root] {
		root = parent[root]
	}
	return root
}

// add u->u into union
func join(u, v int) {
	//u = find_1(u)
	//v = find_1(v)
	//if u == v {
	//	return
	//}
	//parent[v] = u

	if find_1(u) != find_1(v) {
		// build relationship between roots of u and v
		parent[find_1(v)] = find_1(u)
	}
}

func isSameRoot(u, v int) bool {
	//u = find_1(u)
	//v = find_1(v)
	//return u == v
	return find_1(u) == find_1(v)
}

func findRedundantConnection_1(edges [][]int) []int {
	initialize()
	for i := 0; i < len(edges); i++ {
		if isSameRoot(edges[i][0], edges[i][1]) {
			return edges[i]
		} else {
			join(edges[i][0], edges[i][1])
		}
	}
	return []int{}
}
