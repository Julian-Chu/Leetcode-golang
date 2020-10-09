package leetcode959

func regionsBySlashes(grid []string) int {
	m := len(grid)
	size := m * m * 4
	u := newUnion(size)

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			baseIndex := (i*m + j) * 4
			top := baseIndex + 0
			right := baseIndex + 1
			down := baseIndex + 2
			left := baseIndex + 3
			switch grid[i][j] {
			case '\\':
				u.union(top, right)
				u.union(down, left)
			case '/':
				u.union(top, left)
				u.union(down, right)
			default:
				u.union(top, right)
				u.union(right, down)
				u.union(down, left)
			}

			if j+1 < m {
				rsl := baseIndex + 4 + 3
				u.union(right, rsl)
			}

			if i+1 < m {
				dst := baseIndex + 4*m
				u.union(down, dst)
			}
		}
	}
	return u.count
}

type union struct {
	parent []int
	count  int
}

func newUnion(size int) *union {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &union{
		parent: parent,
		count:  size,
	}
}

func (u union) find(i int) int {
	if u.parent[i] == i {
		return i
	}

	return u.find(u.parent[i])
}

func (u *union) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
	u.count--
}
