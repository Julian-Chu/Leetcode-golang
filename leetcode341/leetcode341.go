package leetcode341

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

type NestedIterator struct {
	idx  int
	nums []int
	//listIdx int
	//List []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	nums := make([]int, 0, len(nestedList))
	for _, item := range nestedList {
		flatten(item, &nums)
	}
	return &NestedIterator{
		idx:  -1,
		nums: nums,
	}
}

func (this *NestedIterator) Next() int {
	this.idx++
	return this.nums[this.idx]
}

func (this *NestedIterator) HasNext() bool {
	return this.idx < len(this.nums)-1
}

func flatten(item *NestedInteger, nums *[]int) {
	if item.IsInteger() {
		*nums = append(*nums, item.GetInteger())
	} else {
		list := item.GetList()
		for _, nestedInteger := range list {
			if nestedInteger.IsInteger() {
				*nums = append(*nums, nestedInteger.GetInteger())
			} else {
				flatten(nestedInteger, nums)
			}
		}
	}
}
