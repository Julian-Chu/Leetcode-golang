package leetcode526

func countArrangement(N int) int {
	a := make([]int, N+1)

	for i := 1; i < N+1; i++ {
		a[i] = i
	}

	cnt := 0
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == 0 {
			cnt++
			return
		}

		for i := idx; i > 0; i-- {
			a[idx], a[i] = a[i], a[idx]
			if isBeautiful(a[idx], idx) {
				dfs(idx - 1)
			}
			a[idx], a[i] = a[i], a[idx]
		}
	}

	dfs(N)
	return cnt
}

func isBeautiful(i, j int) bool {
	return i%j == 0 || j%i == 0
}
