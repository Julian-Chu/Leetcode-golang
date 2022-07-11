package LeetCode_692_TopKFrequentWords

import (
	"container/heap"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)

	for i := range words {
		m[words[i]]++
	}
	h := &maxHeap{}
	for key, v := range m {
		heap.Push(h, &Item{
			Word:  key,
			Count: v,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]string, k)

	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(*Item).Word
	}

	return res
}

type Item struct {
	Word  string
	Count int
}

type maxHeap []*Item

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m maxHeap) Less(i, j int) bool {
	if m[i].Count == m[j].Count {
		return m[i].Word > m[j].Word
	}

	return m[i].Count < m[j].Count
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(*Item))
}

func (m *maxHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return res
}

func topKFrequent_1(words []string, k int) []string {
	m := make(map[string]int)

	for _, w := range words {
		m[w]++
	}

	set := make([]string, 0, len(m))
	for k := range m {
		set = append(set, k)
	}
	sort.Slice(set, func(i, j int) bool {
		return m[set[i]] > m[set[j]] || m[set[i]] == m[set[j]] && set[i] < set[j]
	})

	return set[:k]
}
