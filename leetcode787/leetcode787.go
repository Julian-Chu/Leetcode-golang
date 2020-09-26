package leetcode787

import "container/heap"

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	fmap := make([][][]int, n)
	for i := range flights {
		src := flights[i][0]
		fmap[src] = append(fmap[src], flights[i][1:])
	}

	pq := make(PQ, 0, n)

	heap.Push(&pq, &city{
		price:     0,
		id:        src,
		countdown: K + 1,
	})

	for len(pq) > 0 {
		ct, _ := heap.Pop(&pq).(*city)
		if ct.id == dst {
			return ct.price
		}

		if ct.countdown > 0 {
			nexts := fmap[ct.id]
			for _, n := range nexts {
				heap.Push(&pq, &city{
					id:        n[0],
					price:     ct.price + n[1],
					countdown: ct.countdown - 1,
				})
			}
		}
	}
	return -1
}

type city struct {
	id, price, countdown int
}

type PQ []*city

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].price < pq[j].price
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*city))
}

func (pq *PQ) Pop() interface{} {
	res := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return res
}
