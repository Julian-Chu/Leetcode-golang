package LeetCode_685_RedundantConnectionII

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

var (
	n      = 1001
	parent = make([]int, n)
)

func initialize() {
	for i := range parent {
		parent[i] = i
	}
}

func find(u int) int {
	root := u
	for root != parent[root] {
		root = parent[root]
	}
	return root
}

func isSameRoot(u, v int) bool {
	return find(u) == find(v)
}

func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	parent[v] = u
}

func isTreeAfterRemoveEdge(edges [][]int, deleteEdge int) bool {
	initialize()
	for i := 0; i < len(edges); i++ {
		if i == deleteEdge {
			continue
		}
		if isSameRoot(edges[i][0], edges[i][1]) {
			return false
		}
		join(edges[i][0], edges[i][1])
	}
	return true
}

func getRemoveEdge(edges [][]int) []int {
	initialize()
	for i := 0; i < len(edges); i++ {
		if isSameRoot(edges[i][0], edges[i][1]) {
			return edges[i]
		}
		join(edges[i][0], edges[i][1])
	}
	return []int{}
}

func findRedundantDirectedConnection_1(edges [][]int) []int {
	inDegree := make([]int, len(parent))
	for i := range edges {
		inDegree[edges[i][1]] += 1
	}
	twoDegree := make([]int, 0)
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			twoDegree = append(twoDegree, i)
		}
	}

	if len(twoDegree) > 0 {
		if isTreeAfterRemoveEdge(edges, twoDegree[0]) {
			return edges[twoDegree[0]]
		}
		return edges[twoDegree[1]]
	}

	return getRemoveEdge(edges)
}
