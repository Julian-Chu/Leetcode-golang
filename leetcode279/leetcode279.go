package leetcode279

func numSquares(n int) int {
	sn := make([]int, 0)

	for i := 1; i < n; i++ {
		v := i * i

		if v > n {
			break
		}
		sn = append(sn, v)
	}
	minCnt := n
	var recur func(int, int, []int)
	recur = func(rest, cycles int, curRes []int) {
		if cycles >= minCnt {
			return
		}
		if rest == 0 {
			if len(curRes) < minCnt {
				minCnt = len(curRes)
			}
			return
		}
		curRes = curRes[:len(curRes):len(curRes)]
		for i := len(sn) - 1; i >= 0; i-- {
			if sn[i] > rest {
				continue
			}

			recur(rest-sn[i], cycles+1, append(curRes, sn[i]))
		}
	}

	recur(n, 0, []int{})
	return minCnt
}
