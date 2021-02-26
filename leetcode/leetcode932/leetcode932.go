package leetcode932

func beautifulArray(N int) []int {
	res := []int{1}
	for len(res) < N {
		tmp := make([]int, 0, len(res)*2)
		for _, v := range res {
			if v*2-1 <= N {
				tmp = append(tmp, v*2-1)
			}
		}

		for _, v := range res {
			if v*2 <= N {
				tmp = append(tmp, v*2)
			}
		}
		res = tmp
	}
	return res
}
