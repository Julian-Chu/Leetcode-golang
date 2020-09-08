package leetcode373

import (
	"sort"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	l1 := k
	if len(nums1) < k {
		l1 = len(nums1)
	}
	l2 := k
	if len(nums2) < k {
		l2 = len(nums2)
	}
	nums1 = nums1[:l1]
	nums2 = nums2[:l2]

	pairs := make([][]int, 0, len(nums1)*len(nums2))

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			pairs = append(pairs, []int{nums1[i], nums2[j]})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0]+pairs[i][1] < pairs[j][0]+pairs[j][1]
	})

	if len(pairs) < k {
		return pairs
	}
	return pairs[:k]
}
