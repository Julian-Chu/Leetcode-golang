package monotonicqueue

// MonotonicQueue A monotonic Queue is a data structure the elements from the front to the end is
//strictly either increasing or decreasing
type MonotonicQueue struct {
	queue []int
}

func NewMyQueue() *MonotonicQueue {
	return &MonotonicQueue{
		queue: make([]int, 0),
	}
}

func (m *MonotonicQueue) Front() int {
	return m.queue[0]
}

func (m *MonotonicQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MonotonicQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MonotonicQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MonotonicQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}
