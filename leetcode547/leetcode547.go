package leetcode547

func findCircleNum(m [][]int) int {
	n := len(m)

	friends := make([]int, n)
	res := n
	for i := 0; i < n; i++ {
		friends[i] = i
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if m[i][j] == 1 {
				if friends[i] != friends[j] {
					res--
					prev := friends[i]
					next := friends[j]
					for k := range friends {
						if friends[k] == prev {
							friends[k] = next
						}
					}
				}
			}
		}
	}

	return res
}
