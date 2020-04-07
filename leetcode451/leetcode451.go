package leetcode451

import "sort"

func frequencySort(s string) string {
	type pair struct {
		ch  rune
		cnt int
	}

	if len(s) <= 2 {
		return s
	}

	m := make(map[rune]*pair)

	for _, ch := range s {
		if _, ok := m[ch]; !ok {
			m[ch] = &pair{
				ch:  ch,
				cnt: 1,
			}
			continue
		}

		m[ch].cnt++
	}

	pairs := make([]*pair, 0, len(m))

	for _, p := range m {
		pairs = append(pairs, p)
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].cnt > pairs[j].cnt {
			return true
		}
		return false
	})

	res := ""

	for _, p := range pairs {
		for i := 0; i < p.cnt; i++ {
			res += string(p.ch)
		}
	}

	return res
}
