package leetcode787

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := make([][]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for _, flight := range flights {
		src := flight[0]
		dst := flight[1]
		price := flight[2]
		graph[src][dst] = price
	}

	var dfs func(src, k, price int) int

	maxPrice := 100 * 10000
	dfs = func(src, k, price int) int {
		if src == dst {
			return price
		}
		if k == 0 {
			return maxPrice
		}
		min := maxPrice
		for d := 0; d < n; d++ {
			if graph[src][d] < 1 {
				continue
			}
			if visited[d] {
				continue
			}
			visited[src] = true
			sum := price + graph[src][d]
			sum = dfs(d, k-1, sum)
			if sum < min {
				min = sum
			}
			visited[d] = false
		}
		return min
	}

	price := dfs(src, K+1, 0)
	if price >= maxPrice {
		return -1
	}
	return price
}
