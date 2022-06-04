package minheap

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}
	h := &minHeap{}
	heap.Init(h)
	for key, v := range m {
		heap.Push(h, [2]int{key, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

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
