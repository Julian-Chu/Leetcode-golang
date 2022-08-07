package solution3

type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			v := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, v)
		}
	}
	top := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return top
}

func (this *MyQueue) Peek() int {
	v := this.Pop()
	this.out = append(this.out, v)
	return v
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
