package leetcode1202

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	dsu := make(Dsu, len(s))
	dsu.init()
	for _, pair := range pairs {
		dsu.Union(pair[0], pair[1])
	}

	m := make(map[int][]byte)

	for i := range s {
		b := s[i]
		p := dsu.Parent(i)
		m[p] = append(m[p], b)
	}

	for k, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		m[k] = v
	}

	ret := make([]byte, len(s))
	for i := range ret {
		p := dsu.Parent(i)
		group := m[p]
		ret[i] = group[0]
		m[p] = group[1:]
	}
	return string(ret)

}

type Dsu []int

func (d Dsu) init() {
	for i := range d {
		d[i] = i
	}
}

func (d Dsu) Parent(x int) int {
	for x != d[x] {
		x, d[x] = d[x], d[d[x]]
	}
	return x
}

func (d Dsu) Union(x, y int) {
	xp, yp := d.Parent(x), d.Parent(y)
	if xp == yp {
		return
	}
	d[xp] = yp
}
