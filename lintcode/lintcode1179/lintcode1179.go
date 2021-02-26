package lintcode1179

/**
 * @param M: a matrix
 * @return: the total number of friend circles among all the students
 */
func findCircleNum(M [][]int) int {
	if M == nil {
		return 0
	}
	n := len(M)

	circles := 0
	for i := 0; i < n; i++ {
		if M[i][i] != 1 {
			continue
		}
		circles++
		queue := []int{i}

		for len(queue) > 0 {
			next_q := []int{}
			for _, friend := range queue {
				M[friend][friend] = 0
				for k := 0; k < n; k++ {
					if M[friend][k] == 1 && M[k][k] == 1 {
						next_q = append(next_q, k)
					}
				}
				queue = next_q
			}
		}
	}
	return circles
}
