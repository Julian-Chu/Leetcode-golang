package solutionOneQueue

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	if len(this.queue) == 0 {
		return 0
	}

	n := len(this.queue) - 1
	for n > 0 {
		n--
		v := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, v)
	}
	v := this.queue[0]
	this.queue = this.queue[1:]

	return v
}

func (this *MyStack) Top() int {
	v := this.Pop()
	this.Push(v)
	return v
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
