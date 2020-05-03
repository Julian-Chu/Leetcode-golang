package leetcode621

import "sort"

func leastInterval(tasks []byte, n int) int {
	ts := make([]int, 26)
	for _, v := range tasks {
		ts[v-'A']++
	}

	sort.Ints(ts)

	i := 25 // last sequence
	// A B C D   => if n > (count of a line) ,
	// A B C D
	// A B C D
	// A B        last sequence

	for _, v := range ts {
		if v == ts[25] {
			i--
		}
	}

	res := (ts[25]-1)*(n+1) + 25 - i

	if res > len(tasks) {
		return res
	}

	return len(tasks)
}
