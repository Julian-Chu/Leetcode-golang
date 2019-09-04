package leetcode646

import (
	"sort"
)

func findLongestChain(pairs [][]int) int {
	n1 := len(pairs)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})

	res := 0
	previousHead := 1 << 32
	for i := n1 - 1; i >= 0; i-- {
		nextTail := pairs[i][1]
		if previousHead > nextTail {
			res++
			previousHead = pairs[i][0]
		}
	}
	return res
}

//func findLongestChain(pairs [][]int) int {
//	sort.Slice(pairs, func(i, j int) bool {
//		if pairs[i][0] == pairs[j][0] {
//			return pairs[i][1] < pairs[j][1]
//		}
//		return pairs[i][1] < pairs[j][1]
//	})
//
//	res := make([]int, n1+1)
//	maxLen := 0
//	for i := 0; i < n1; i++ {
//		for j := 0; j < i; j++ {
//			if pairs[i][0] > pairs[j][1] {
//				res[i] = max(res[i], res[j]+1)
//			}
//		}
//		if maxLen < res[i]+1 {
//			maxLen = res[i] + 1
//		}
//	}
//	return maxLen
//}
//
//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
