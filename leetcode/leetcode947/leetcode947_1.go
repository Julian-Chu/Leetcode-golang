package leetcode947

func removeStones1(stones [][]int) int {
	roots := make(map[int]int)
	var rx, ry int

	for _, stone := range stones {
		rx = getRoot(stone[0], roots)
		ry = getRoot(stone[1]+10000, roots)
		roots[ry] = rx
	}

	cnt := 0
	for k, v := range roots {
		if k == v {
			cnt++
		}
	}
	return len(stones) - cnt
}

func getRoot(node int, roots map[int]int) int {
	if _, ok := roots[node]; !ok {
		roots[node] = node
		return node
	}

	if roots[node] == node {
		return node
	}

	return getRoot(roots[node], roots)
}
