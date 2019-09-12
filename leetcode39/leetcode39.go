package leetcode39

func combinationSum(candidates []int, target int) [][]int {
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
			temp := make([]int, len(cur))
			copy(temp, cur)
			temp = append(temp, candidates[i])
			dfs(temp, target-candidates[i], i)
		}
	}

	dfs([]int{}, target, 0)

	return res
}
