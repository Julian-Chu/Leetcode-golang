package leetcode210

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjacency := make([][]int, numCourses)

	for _, pre := range prerequisites {
		adjacency[pre[1]] = append(adjacency[pre[1]], pre[0])
	}

	orderings := make([]int, 0, numCourses)
	flags := make([]int, numCourses)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		switch flags[idx] {
		case -1:
			return true
		case 1:
			return false
		}

		flags[idx] = 1
		for _, i := range adjacency[idx] {
			if !dfs(i) {
				return false
			}
		}
		flags[idx] = -1
		orderings = append(orderings, idx)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}
	l, r := 0, numCourses-1

	for l < r {
		orderings[l], orderings[r] = orderings[r], orderings[l]
		l++
		r--
	}
	return orderings
}
