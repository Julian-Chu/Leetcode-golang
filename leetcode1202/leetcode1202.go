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

	m := make(map[int][]byte)

	bs := []byte(s)
	for i, b := range bs {
		p := getParent(i, parent)
		m[p] = append(m[p], b)
	}

	for k, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
		m[k] = v
	}

	res := make([]byte, n)

	for i := range res {
		p := getParent(i, parent)
		v := m[p][len(m[p])-1]
		m[p] = m[p][:len(m[p])-1]
		res[i] = v
	}
	return string(res)
}

func getParent(i int, parent []int) int {
	if i == parent[i] {
		return i
	}

	return getParent(parent[i], parent)
}
