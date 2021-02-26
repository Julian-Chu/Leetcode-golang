package leetcode947

func removeStones(stones [][]int) int {
	union := newUnion(20000)

	for _, s := range stones {
		union.Union(s[0], s[1]+10000)
	}

	keep := make(map[int]int)

	for _, s := range stones {
		root := union.find(s[0])
		keep[root]++
	}

	return len(stones) - len(keep)
}

type union struct {
	parent []int
}

func newUnion(size int) *union {
	parent := make([]int, size)

	for i := range parent {
		parent[i] = i
	}

	return &union{parent: parent}
}

func (u union) find(i int) int {
	if u.parent[i] == i {
		return i
	}

	return u.find(u.parent[i])
}

func (u *union) Union(x, y int) {
	xp, yp := u.find(x), u.find(y)

	u.parent[xp] = yp
}
