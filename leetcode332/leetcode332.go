package leetcode332

import "sort"

func findItinerary(tickets [][]string) []string {
	nexts := make(map[string][]string, len(tickets))
	res := make([]string, 0, len(tickets))

	for _, ticket := range tickets {
		nexts[ticket[0]] = append(nexts[ticket[0]], ticket[1])
	}

	for _, next := range nexts {
		sort.Strings(next)
	}
	var dfs func(from string)
	dfs = func(from string) {
		for len(nexts[from]) > 0 {
			next := nexts[from][0]
			nexts[from] = nexts[from][1:]
			dfs(next)
		}
		res = append(res, from)
	}

	dfs("JFK")

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
