package leetcode703

import "container/heap"

type KthLargest struct {
	nums MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := MinHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}
	return KthLargest{
		nums: h,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.nums, val)
	if this.nums.Len() > this.k {
		heap.Pop(&this.nums)
	}

	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}
