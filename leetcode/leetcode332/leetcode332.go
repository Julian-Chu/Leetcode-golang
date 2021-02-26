package leetcode332

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
