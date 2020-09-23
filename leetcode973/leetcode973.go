package leetcode973

import "container/heap"

func kClosest(points [][]int, K int) [][]int {
	h := minHeap(points)
	heap.Init(&h)
	for h.Len() > K {
		heap.Pop(&h)
	}
	return h
}

type minHeap [][]int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i][0]*m[i][0]+m[i][1]*m[i][1] > m[j][0]*m[j][0]+m[j][1]*m[j][1]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	(*m) = append((*m), x.([]int))
}

func (m *minHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}
