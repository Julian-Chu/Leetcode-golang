package leetcode207

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return false
	}

	preq := make([]map[int]bool, numCourses)
	for i := range preq {
		preq[i] = make(map[int]bool)
	}

	for i := range prerequisites {
		m := preq[prerequisites[i][1]]
		if _, ok := m[prerequisites[i][0]]; !ok {
			m[prerequisites[i][0]] = true
		}
	}

	q := make([]int, 0, 100000)

	for i := range preq {
		if len(preq[i]) == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		qNext := make([]int, 0, 100000)
		for _, v := range q {
			for i, m := range preq {

				if len(m) == 0 {
					continue
				}
				if _, ok := m[v]; ok {
					delete(m, v)
				}
				if len(m) == 0 {
					qNext = append(qNext, i)
				}
			}
		}
		q = qNext
	}

	for _, m := range preq {
		if len(m) > 0 {
			return false
		}
	}
	return true

}
