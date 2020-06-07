package leetcode225

type MyStack struct {
	queue Queue
	top   int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: Queue{slice: make([]int, 0)},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.enqueue(x)
	this.top = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.queue.isEmpty() {
		return 0
	}

	nextq := Queue{}

	for this.queue.size() != 1 {
		nextq.enqueue(this.queue.dequeue())
	}
	res := this.queue.dequeue()
	this.queue = nextq
	if this.queue.isEmpty() {
		this.top = 0
	} else {
		this.top = nextq.slice[0]
	}
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue.slice) == 0
}

type Queue struct {
	slice []int
}

func (q *Queue) dequeue() int {
	queue := q.slice
	res := queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	q.slice = queue
	return res
}

func (q *Queue) enqueue(x int) {
	queue := q.slice
	nextq := make([]int, len(queue)+1)
	copy(nextq[1:], queue)
	nextq[0] = x
	q.slice = nextq
}

func (q Queue) size() int {
	return len(q.slice)
}

func (q Queue) isEmpty() bool {
	return q.size() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
