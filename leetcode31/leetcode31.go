package leetcode31

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	end := len(nums) - 1
	for i := end; i >= 0; i-- {
		if i > 0 && nums[i] > nums[i-1] {
			if i == end {
				nums[end], nums[end-1] = nums[end-1], nums[end]
			} else {
				getNextPremutation(nums, i)
			}
			break
		} else if i == 0 {
			sort.Ints(nums)
		}
	}
	fmt.Println(nums)
}

func getNextPremutation(nums []int, i int) {
	end := len(nums) - 1
	idx := i
	for j := end; j >= i; j-- {
		if nums[j] > nums[i-1] {
			idx = j
			break
		}
	}
	nums[i-1], nums[idx] = nums[idx], nums[i-1]
	l, r := i, end
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
