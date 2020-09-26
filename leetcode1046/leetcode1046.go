package leetcode1046

import "container/heap"

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	h := StonesMaxHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		s1 := heap.Pop(&h).(int)
		s2 := heap.Pop(&h).(int)
		heap.Push(&h, s1-s2)
	}

	return h[0]
}

type StonesMaxHeap []int

func (s StonesMaxHeap) Len() int {
	return len(s)
}

func (s StonesMaxHeap) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s StonesMaxHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *StonesMaxHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *StonesMaxHeap) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
