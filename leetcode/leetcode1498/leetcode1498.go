package leetcode1498

import (
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	mod := int(1e9 + 7)
	res := 0

	pow2 := make([]int, len(nums))
	pow2[0] = 1
	for i := 1; i < len(nums); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			res += pow2[r-l]
			l++
		}
	}
	return res % mod
}
