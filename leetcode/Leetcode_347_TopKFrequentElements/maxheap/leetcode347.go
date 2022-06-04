package maxheap

import (
	"container/heap"
)

// min heap is better than max heap, no need to push all record to heap, maintain k size during
// pop and push

type Record struct {
	val  int
	freq int
}

type Heap struct {
	elems []Record
}

func (h *Heap) Len() int {
	return len(h.elems)
}

func (h *Heap) Less(i, j int) bool {
	return h.elems[i].freq > h.elems[j].freq
}

func (h *Heap) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *Heap) Push(x interface{}) {
	h.elems = append(h.elems, x.(Record))
}

func (h *Heap) Pop() interface{} {
	n := h.Len()
	x := h.elems[n-1]
	h.elems = h.elems[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	records := make([]Record, 0)

	for val, count := range freq {
		records = append(records, Record{val: val, freq: count})
	}
	h := &Heap{}
	heap.Init(h)

	for _, v := range records {
		heap.Push(h, v)
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(Record).val)
	}
	return res
}
