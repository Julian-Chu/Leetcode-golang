package leetcode341

import "Leetcode-golang/utils"

type NestedInteger = utils.NestedInteger

type NestedIterator struct {
	idx  int
	nums []int
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
