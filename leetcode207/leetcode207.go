package leetcode207

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return false
	}
	paths := make([][]int, 0)

	var dfs func(cur [2]int, m *map[[2]int]bool)
	dfs = func(cur [2]int, m *map[[2]int]bool) {
		if _, ok := (*m)[cur]; ok {
			return
		}

	}
	return res
}
