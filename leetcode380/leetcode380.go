package leetcode380

import "math/rand"

type RandomizedSet struct {
	idx  map[int]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{idx: make(map[int]int), nums: make([]int, 0)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.idx[val] = len(this.nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.idx[val]; ok {
		size := len(this.nums)
		this.nums[i], this.nums[size-1] = this.nums[size-1], this.nums[i]
		this.idx[this.nums[i]] = i
		this.nums = this.nums[:size-1]
		// move delete to last due to case that remove last element
		delete(this.idx, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
