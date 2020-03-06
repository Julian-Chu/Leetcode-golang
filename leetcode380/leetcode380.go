package leetcode380

import "math/rand"

type RandomizedSet struct {
	numMapper map[int]bool
	nums      []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{numMapper: make(map[int]bool), nums: make([]int, 0)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numMapper[val]; !ok {
		this.numMapper[val] = true
		this.nums = append(this.nums, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numMapper[val]; ok {
		delete(this.numMapper, val)
		for i := range this.nums {
			if this.nums[i] == val {
				size := len(this.nums)
				this.nums[i], this.nums[size-1] = this.nums[size-1], this.nums[i]
				this.nums = this.nums[: size-1 : size-1]
				break
			}
		}

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
