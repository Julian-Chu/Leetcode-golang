package leetcode216

func combinationSum3(k int, n int) [][]int {
	if k == 0 || n < 3 {
		return [][]int{}
	}

	q := make([]map[int]bool, 0, 9)
	res := make([][]int, 0)
	for i := 0; i < 9; i++ {
		q = append(q, map[int]bool{i + 1: true})
	}

	for len(q) > 0 {
		if len(q[0]) == k {
			break
		}
		next := make([]map[int]bool, 0)

		for idx := range q {
			sum := 0
			max := 0
			for v := range q[idx] {
				sum += v
				if v > max {
					max = v
				}
			}
			upper := 9
			if n-sum < 9 {
				upper = n - sum
			}
			for i := 0; i < upper; i++ {
				if i+1 <= max {
					continue
				}
				if _, ok := q[idx][i+1]; ok {
					continue
				}
				newMap := make(map[int]bool)
				for k, v := range q[idx] {
					newMap[k] = v
				}
				newMap[i+1] = true
				next = append(next, newMap)
			}
		}
		q = next
	}

	for i := range q {
		s := make([]int, 0, k)
		sum := 0
		for v := range q[i] {
			s = append(s, v)
			sum += v
		}
		if sum != n {
			continue
		}
		res = append(res, s)
	}

	return res
}
