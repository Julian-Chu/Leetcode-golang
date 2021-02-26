package leetcode78

import "sort"

func subsets_1(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{{}}

	var dfs func([]int, int)
	dfs = func(cur []int, idx int) {
		cur = cur[:len(cur):len(cur)]
		cur = append(cur, nums[idx])
		res = append(res, cur)

		for i := range nums {
			if i <= idx {
				continue
			}
			dfs(cur, i)
		}
	}

	for i := range nums {
		dfs([]int{}, i)
	}
	return res
}
