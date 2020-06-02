package leetcode457

import "unsafe"

func circularArrayLoop(nums []int) bool {
	size := len(nums)

	for i := 0; i < size; i++ {
		nums[i] %= size
	}

	bits := uint(unsafe.Sizeof(size) - 1)

	for i, n := range nums {
		mark := (i + size) * (n>>bits | 1)

		for -size < n && n < size && n != 0 {
			nums[i] = mark

			i = (n + i + size) % size
			n = nums[i]

			if n == mark {
				return true
			}

			if n*mark < 0 {
				break
			}
		}
	}
	return false
}
