package leetcode526

func countArrangement(N int) int {
	arr := make([]bool, N+1)

	for i := 1; i < N+1; i++ {
		arr[i] = true
	}

	cnt := 0
	var dfs func(restNums []bool, idx int)
	dfs = func(restNums []bool, idx int) {
		if idx == N+1 {
			cnt++
			return
		}
		for i, v := range restNums {
			if v == false {
				continue
			}
			if i%idx == 0 || idx%i == 0 {
				rest := append([]bool{}, restNums...)
				rest[i] = false
				dfs(rest, idx+1)
			}
		}
	}

	dfs(arr, 1)
	return cnt
}
