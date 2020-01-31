package leetcode155

import "math"

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), min: math.MaxInt64}
}

func (this *MinStack) Push(x int) {
	if x < this.min {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	res := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if res == this.min {
		min := math.MaxInt64
		for _, v := range this.stack {
			if v < min {
				min = v
			}
		}
		this.min = min
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
