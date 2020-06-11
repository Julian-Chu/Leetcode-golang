package utils

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Num  int
	List []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested List.
func (this NestedInteger) IsInteger() bool {
	return len(this.List) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested List
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.Num
}

// Set this NestedInteger to hold a single integer.
func (this *NestedInteger) SetInteger(value int) {
	this.Num = value
}

// Set this NestedInteger to hold a nested List and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	this.List = append(this.List, &elem)
}

// Return the nested List that this NestedInteger holds, if it holds a nested List
// The List length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.List
}
