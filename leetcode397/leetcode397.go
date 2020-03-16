package leetcode397

func integerReplacement(n int) int {
	if n <= 1 {
		return 0
	}

	powerOfTwo := make(map[int]int)

	cnt := 0

	for i := 1; i <= n+1; i *= 2 {
		powerOfTwo[i] = cnt
		cnt++
	}

	min := 1 << 31
	var dfs func(int, int)
	dfs = func(v int, curCnt int) {
		if curCnt > min {
			return
		}
		if cnt, ok := powerOfTwo[v]; ok {
			if min > (cnt + curCnt) {
				min = cnt + curCnt
			}
			return
		}
		if v%2 == 0 {
			dfs(v/2, curCnt+1)
		} else {
			dfs(v+1, curCnt+1)
			dfs(v-1, curCnt+1)
		}
	}

	dfs(n, 0)

	return min
}
