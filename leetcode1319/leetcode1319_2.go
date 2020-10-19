package leetcode1319

func makeConnected_1(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	u := NewDSU(n)
	cnt := 0
	for _, c := range connections {
		if u.union(c[0], c[1]) {
			cnt++
		}
	}
	return n - 1 - cnt
}

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DSU{
		parent: parent,
		rank:   rank,
	}
}

func (d DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d DSU) union(x int, y int) bool {
	xp := d.find(x)
	yp := d.find(y)

	if xp == yp {
		return false
	} else if d.rank[xp] > d.rank[yp] {
		d.parent[xp] = yp
	} else if d.rank[xp] > d.rank[yp] {
		d.parent[yp] = xp
	} else {
		d.parent[yp] = xp
		d.rank[xp]++
	}
	return true
}
