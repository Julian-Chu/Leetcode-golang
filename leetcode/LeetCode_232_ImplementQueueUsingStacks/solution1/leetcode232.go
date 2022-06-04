package solution1

type MyQueue struct {
	in, out Stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  Stack{},
		out: Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Pop()
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Peek()
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

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *Stack) Peek() int {
	x := (*s)[len(*s)-1]
	return x
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
