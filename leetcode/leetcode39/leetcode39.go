package leetcode39

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if len(candidates) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)

	var dfs func([]int, int, int)
	dfs = func(cur []int, target int, index int) {

		if target < 0 {
			return
		}
		if target == 0 {
			res = append(res, cur)
			return
		}

		for i := index; i < len(candidates); i++ {
			if candidates[i] > target { //avoid unnecessary stack
				break
			}
			//temp := make([]int, len(cur))
			//copy(temp, cur)
			temp := cur[:len(cur):len(cur)] //avoid unnecessary copy
			temp = append(temp, candidates[i])
			dfs(temp, target-candidates[i], i)
		}
	}

	dfs([]int{}, target, 0)

	return res
}

func combinationSum_1(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var dfs func(subset []int, start, sum int)
	dfs = func(subset []int, start, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, subset...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				return
			}

			subset = append(subset, candidates[i])
			dfs(subset, i, sum+candidates[i])
			subset = subset[:len(subset)-1]
		}
	}

	dfs(nil, 0, 0)
	return res
}
