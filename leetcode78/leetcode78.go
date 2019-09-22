package leetcode78

import (
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	res = append(res, []int{})
	if len(nums) == 0 {
		return res
	}

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			return
		}
		newSolutions := make([][]int, 0)
		for j := range res {
			res[j] = res[j][:len(res[j]):len(res[j])]
			next := append(res[j], nums[i])
			newSolutions = append(newSolutions, next)
		}

		res = append(res, newSolutions...)
		dfs(i + 1)
	}
	dfs(0)
	return res
}
