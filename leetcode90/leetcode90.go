package leetcode90

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	var dfs func(int, []int)

	dfs = func(index int, solution []int) {
		solution = solution[:len(solution):len(solution)]
		res = append(res, solution)
		for i := index; i < len(nums); i++ {
			if i == index || nums[i] != nums[i-1] {
				dfs(i+1, append(solution, nums[i]))
			}
		}
	}

	dfs(0, []int{})
	return res
}
