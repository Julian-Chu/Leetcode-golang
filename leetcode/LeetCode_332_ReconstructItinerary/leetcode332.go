package LeetCode_332_ReconstructItinerary

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	tickets_map := make(map[string][]string)
	tickets_map_visited := make(map[string][]bool)

	for _, t := range tickets {
		tickets_map[t[0]] = append(tickets_map[t[0]], t[1])
		tickets_map_visited[t[0]] = append(tickets_map_visited[t[0]], false)
	}

	for _, v := range tickets_map {
		sort.Strings(v)
	}
	cur := "JFK"
	travel := []string{"JFK"}
	dfs(cur, len(tickets), tickets_map, tickets_map_visited, &travel)
	return travel
}

func dfs(cur string, left int, hash map[string][]string, hash2 map[string][]bool, travel *[]string) bool {
	if left == 0 {
		return true
	}
	if len(hash[cur]) == 0 {
		return false
	}

	for i, v := range hash2[cur] {
		if v == false {
			hash2[cur][i] = true
			*travel = append(*travel, hash[cur][i])
			if dfs(hash[cur][i], left-1, hash, hash2, travel) == true {
				return true
			}
			hash2[cur][i] = false
			*travel = (*travel)[:len(*travel)-1]
		}
	}
	return false
}

func findItinerary_1(tickets [][]string) []string {
	if len(tickets) == 0 {
		return nil
	}

	nexts := make(map[string][]string)
	nextsUsed := make(map[string][]bool)

	for _, ticket := range tickets {
		from := ticket[0]
		to := ticket[1]
		if _, ok := nexts[from]; !ok {
			nexts[from] = make([]string, 0)
		}
		nexts[from] = append(nexts[from], to)
	}

	for from, dests := range nexts {
		sort.Strings(dests)
		nextsUsed[from] = make([]bool, len(dests))
	}

	var res []string
	restTickets := len(tickets)
	var dfs func(tmp []string, restTickers int) bool
	dfs = func(tmp []string, restTickets int) bool {
		if restTickets == 0 {
			res = append([]string{}, tmp...)
			return true
		}

		from := tmp[len(tmp)-1]

		for i := 0; i < len(nexts[from]); i++ {
			if nextsUsed[from][i] {
				continue
			}

			nextsUsed[from][i] = true
			tmp = append(tmp, nexts[from][i])
			if dfs(tmp, restTickets-1) {
				return true
			}
			tmp = tmp[:len(tmp)-1]
			nextsUsed[from][i] = false
		}
		return false
	}

	dfs([]string{"JFK"}, restTickets)
	return res
}

func findItinerary_2(tickets [][]string) []string {
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
