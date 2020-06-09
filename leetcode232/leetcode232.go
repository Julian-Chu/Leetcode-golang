package leetcode232

type MyQueue struct {
	a, b Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		a: Stack{},
		b: Stack{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.a.IsEmpty() {
		this.a, this.b = this.b, this.a
	}

	for !this.a.IsEmpty() {
		this.b.Push(this.a.Pop())
	}
	this.a.Push(x)

	for !this.b.IsEmpty() {
		this.a.Push(this.b.Pop())
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.a.IsEmpty() && this.b.IsEmpty() {
		return 0
	}
	if this.a.IsEmpty() {
		this.a, this.b = this.b, this.a
	}

	return this.a.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	res := this.a.Pop()
	this.a.Push(res)

	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.a.IsEmpty() && this.b.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
