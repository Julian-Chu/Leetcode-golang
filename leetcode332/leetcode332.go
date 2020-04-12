package leetcode332

import "sort"

func findItinerary(tickets [][]string) []string {
	tickets_map := make(map[string][]string)
	for _, arr := range tickets {
		if _, ok := tickets_map[arr[0]]; ok {
			tickets_map[arr[0]] = append(tickets_map[arr[0]], arr[1])
		} else {
			tickets_map[arr[0]] = []string{arr[1]}
		}
	}

	for _, arr := range tickets_map {
		sort.Strings(arr)
	}

	var ret []string
	stack := []string{"JFK"}
	for len(stack) != 0 {
		key := stack[len(stack)-1]
		if len(tickets_map[key]) != 0 {
			var nextkey string
			nextkey, tickets_map[key] = tickets_map[key][0], tickets_map[key][1:]
			stack = append(stack, nextkey)
		} else {
			ret = append([]string{key}, ret...)
			stack = stack[0 : len(stack)-1]
		}
	}
	return ret
}
