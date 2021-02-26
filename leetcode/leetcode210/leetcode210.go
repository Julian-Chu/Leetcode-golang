package leetcode210

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	orderings := make([]int, 0, numCourses)
	for _, prereq := range prerequisites {
		inDegrees[prereq[0]]++
		adjacency[prereq[1]] = append(adjacency[prereq[1]], prereq[0])
	}

	q := make([]int, 0)
	for i, cnt := range inDegrees {
		if cnt == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		orderings = append(orderings, e)
		for _, course := range adjacency[e] {
			inDegrees[course]--
			if inDegrees[course] == 0 {
				q = append(q, course)
			}
		}
	}
	if len(orderings) < numCourses {
		return []int{}
	}
	return orderings
}
