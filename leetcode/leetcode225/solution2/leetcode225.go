package solution2

type MyStack struct {
	in, out []int
}

func Constructor() MyStack {
	return MyStack{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	if len(this.out) == 0 {
		this.out = append(this.out, x)
		return
	}

	e := this.out[0]
	this.in = append(this.in, e)
	this.out[0] = x
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	for len(this.out) == 0 {
		this.out, this.in = this.in, this.out
	}

	for len(this.out) > 1 {
		e := this.out[0]
		this.in = append(this.in, e)
		copy(this.out, this.out[1:])
		this.out = this.out[:len(this.out)-1]
	}

	e := this.out[0]
	this.out = this.out[:0]
	return e
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	for len(this.out) == 0 {
		this.out, this.in = this.in, this.out
	}

	for len(this.out) > 1 {
		e := this.out[0]
		this.in = append(this.in, e)
		copy(this.out, this.out[1:])
		this.out = this.out[:len(this.out)-1]
	}

	return this.out[0]
}

func (this *MyStack) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
