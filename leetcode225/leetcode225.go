package leetcode225

type MyStack struct {
	a, b *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		a: NewQueue(), b: NewQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}
	this.a.Enqueue(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}

	for this.a.Len() > 1 {
		this.b.Enqueue(this.a.Dequeue())
	}
	return this.a.Dequeue()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	res := this.Pop()
	this.Push(res)
	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return (this.a.Len() + this.b.Len()) == 0
}

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Enqueue(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Dequeue() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
