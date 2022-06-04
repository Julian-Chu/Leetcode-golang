package solution1

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	q := &monotonicQueue{
		queue: make([]int, 0),
	}
	for i := range nums {
		if i < k-1 {
			q.Push(nums[i])
			continue
		}

		q.Push(nums[i])
		res = append(res, q.Top())
		if nums[i-k+1] == q.Top() {
			q.Pop()
		}
	}

	return res
}

type monotonicQueue struct {
	queue []int
}

func (q *monotonicQueue) Push(x int) {
	for len(q.queue) > 0 && x > q.queue[len(q.queue)-1] {
		q.queue = q.queue[:len(q.queue)-1]
	}

	q.queue = append(q.queue, x)
}

func (q *monotonicQueue) Top() int {
	return q.queue[0]
}

func (q *monotonicQueue) Pop() int {
	e := q.queue[0]
	//copy(q.queue, q.queue[1:])
	//q.queue = q.queue[:len(q.queue)-1]
	q.queue = q.queue[1:] // better performance
	return e
}
