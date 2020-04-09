package leetcode207

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)

	for _, edge := range prerequisites {
		indegrees[edge[0]]++
		adjacency[edge[1]] = append(adjacency[edge[1]], edge[0])
	}

	q := make([]int, 0, numCourses)

	for i, cnt := range indegrees {
		if cnt == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		course := q[0]
		q = q[1:]

		numCourses--

		for _, cur := range adjacency[course] {
			indegrees[cur]--
			if indegrees[cur] == 0 {
				q = append(q, cur)
			}
		}
	}

	return numCourses == 0
}
