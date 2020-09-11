package leetcode743

func networkDelayTime(times [][]int, N int, K int) int {
	maxInt := int(1e4)

	minTime := make([]int, N+1)
	for i := 1; i <= N; i++ {
		minTime[i] = maxInt
	}

	minTime[K] = 0
	cost := make([][]int, N+1)
	for i := range cost {
		cost[i] = make([]int, N+1)
		for j := 0; j <= N; j++ {
			cost[i][j] = maxInt
		}
	}

	for _, t := range times {
		cost[t[0]][t[1]] = t[2]
	}

	queue := make([]int, 1, N+1)
	queue[0] = K
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 1; v <= N; v++ {
			if minTime[v] > minTime[u]+cost[u][v] {
				minTime[v] = minTime[u] + cost[u][v]
				queue = append(queue, v)
			}
		}
	}

	res := 0
	for i := 1; i <= N; i++ {
		res = max(res, minTime[i])
	}
	if res == maxInt {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
