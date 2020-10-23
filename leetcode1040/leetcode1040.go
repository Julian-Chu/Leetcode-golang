package leetcode1040

import "sort"

// explanation: https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/solution/jie-ti-si-lu-by-owenzzz/

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	return []int{low(stones), high(stones)}
}

func low(s []int) int {
	n := len(s)
	// corner case
	if (s[n-2]-s[0] == n-2 && s[n-1]-s[n-2] > 2) ||
		(s[n-1]-s[1] == n-2 && s[1]-s[0] > 2) {
		return 2
	}
	// sliding window is s[i:j]
	width, i, j := 0, 0, 1
	for ; j < n; j++ {
		if s[j]-s[i] < n {
			continue
		}
		width = max(width, j-i)
		i = j
	}
	width = max(width, j-i)
	return n - width
}

func high(stones []int) int {
	n := len(stones)
	/*
		s1:= stones[n-1]-stones[0]+1-n // space between 2 endpoints
		s2:= min(stones[n-1]-stones[n-2]-1, stones[1]-stones[0]-1) // space can't be moved
		s:= s1-s2 // space can be moved
	*/
	return max(stones[n-2]-stones[0], stones[n-1]-stones[1]) - n + 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
