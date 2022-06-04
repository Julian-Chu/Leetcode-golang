package minheap

import "container/heap"

type minHeap [][2]int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.([2]int))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

func example() {
	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, [2]int{1, 2})
	heap.Pop(h)
}
