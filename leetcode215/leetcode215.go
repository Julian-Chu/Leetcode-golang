package leetcode215

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	temp := maxHeap(nums)
	h := &temp
	heap.Init(h)
	if k == 1 {
		return nums[0]
	}

	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}
	return (*h)[0]
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return res
}
