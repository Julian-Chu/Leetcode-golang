package leetcode1202

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	for _, p := range pairs {
		xp, yp := getParent(p[0], parent), getParent(p[1], parent)
		parent[xp] = yp
	}

	m := make(map[int][]int)

	for i := range s {
		p := getParent(i, parent)
		m[p] = append(m[p], i)
	}

	bs := []byte(s)
	res := make([]byte, n)
	for _, g := range m {
		idx := make([]int, len(g))
		copy(idx, g)
		sort.Slice(idx, func(i, j int) bool {
			return bs[idx[i]] < bs[idx[j]]
		})

		sort.Ints(g)
		for i := range g {
			res[g[i]] = bs[idx[i]]
		}
	}
	return string(res)

}

func getParent(i int, parent []int) int {
	if i == parent[i] {
		return i
	}

	return getParent(parent[i], parent)
}
