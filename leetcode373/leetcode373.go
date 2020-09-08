package leetcode373

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	h := minHeap{}
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(&h, Item{0, nums1[i], nums2[0]})
	}
	var result [][]int
	for len(h) > 0 {
		p := heap.Pop(&h).(Item)
		result = append(result, []int{p.left, p.right})
		k--
		if k == 0 {
			break
		}
		if j := p.idx + 1; j < len(nums2) {
			heap.Push(&h, Item{j, p.left, nums2[j]})
		}
	}
	return result
}

type Item struct {
	idx   int
	left  int
	right int
}

type minHeap []Item

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].left+m[i].right < m[j].left+m[j].right
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(Item))
}

func (m *minHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return res
}
