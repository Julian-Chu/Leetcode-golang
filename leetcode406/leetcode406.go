package leetcode406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		}
		if people[i][0] == people[j][0] && people[i][1] < people[j][1] {
			return true
		}
		return false
	})

	res := make([][]int, len(people))

	for _, p := range people {
		idx := p[1]
		copy(res[idx+1:], res[idx:])
		res[idx] = p
	}

	return res
}
