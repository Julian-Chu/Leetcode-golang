package leetcode1319

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var findParent func(int, []int) int
	findParent = func(i int, parent []int) int {
		if i == parent[i] {
			return i
		}

		return findParent(parent[i], parent)
	}
	cnt := 0
	for _, connection := range connections {
		xp := findParent(connection[0], parent)
		yp := findParent(connection[1], parent)
		if xp == yp {
			continue
		}
		if xp < yp {
			parent[yp] = xp
		} else {
			parent[xp] = yp
		}
		cnt++
	}

	return n - 1 - cnt
}
