package leetcode743

func networkDelayTime(times [][]int, N int, K int) int {
	graphs := make([][]int, N+1)
	values := make([][]int, N+1)
	for _, time := range times {
		graphs[time[0]] = append(graphs[time[0]], time[1])
		values[time[0]] = append(values[time[0]], time[2])
	}

	queue := []int{K}
	visited := make([]int, N+1)
	for i := range visited {
		visited[i] = -1
	}
	visited[K] = 0

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for i, neighbor := range graphs[curr] {
			if visited[neighbor] == -1 || visited[neighbor] > visited[curr]+values[curr][i] {
				visited[neighbor] = visited[curr] + values[curr][i]
				queue = append(queue, neighbor)
			}
		}
	}
	max := -1

	for i := 1; i < N+1; i++ {
		if visited[i] == -1 {
			return -1
		}

		if visited[i] > max {
			max = visited[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func networkDelayTimeHeap(times [][]int, N int, K int) int {
	m := make(map[int]map[int]int)
	for _, time := range times {
		if _, ok := m[time[0]]; !ok {
			m[time[0]] = make(map[int]int)
		}
		m[time[0]][time[1]] = time[2]
	}

	pq := MaxHeap{}
	pq.Push(NextNode{
		Next: K,
		Dist: 0,
	})
	dist := make(map[int]int)
	visited := make(map[int]bool)
	for pq.Len() > 0 {
		cur := pq.Pop().(NextNode)
		curNode := cur.Next
		curDist := cur.Dist

		if _, ok := dist[curNode]; ok {
			if dist[curNode] <= cur.Dist {
				continue
			}
			dist[curNode] = cur.Dist
		} else {
			dist[curNode] = cur.Dist
			visited[curNode] = true
			N--
		}

		if _, ok := m[curNode]; ok {
			nexts := m[curNode]
			for next, dist := range nexts {
				pq.Push(NextNode{
					Dist: curDist + dist,
					Next: next,
				})
			}

		}

	}

	if N > 0 {
		return -1
	}
	res := 0
	for _, d := range dist {
		if d > res {
			res = d
		}
	}
	return res
}

type NextNode struct {
	Next int
	Dist int
}
type MaxHeap []NextNode

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].Dist > m[j].Dist
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(NextNode))
}

func (m *MaxHeap) Pop() interface{} {
	res := (*m)[0]
	*m = (*m)[1:]
	return res
}
