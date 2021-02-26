package leetcode207

func canFinish(numCourses int, prerequisites [][]int) bool {
	flags := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for _, pair := range prerequisites {
		adjacency[pair[1]] = append(adjacency[pair[1]], pair[0])
	}
	var dfs func(int) bool
	dfs = func(i int) bool {
		if flags[i] == -1 {
			return true
		}
		if flags[i] == 1 {
			return false
		}
		flags[i] = 1
		for _, idx := range adjacency[i] {
			if !dfs(idx) {
				return false
			}
		}
		flags[i] = -1
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
