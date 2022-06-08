package LeetCode_77_Combinations

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(n, k, start int, subset []int)
	dfs = func(n, k, start int, subset []int) {
		if len(subset) == k {
			res = append(res, append([]int{}, subset...))
			return
		}

		for i := start; i < n; i++ {
			subset = append(subset, i+1)
			dfs(n, k, i+1, subset)
			subset = subset[:len(subset)-1]
		}
	}

	dfs(n, k, 0, nil)
	return res
}

func combine_1(n int, k int) [][]int {
	nums := []int{}

	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}

	var res [][]int
	var dfs func(nums, subset []int, idx, k int)
	dfs = func(nums, subset []int, start, k int) {
		if k == 0 {
			res = append(res, append([]int{}, subset...))
			return
		}

		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])
			dfs(nums, subset, i+1, k-1)
			subset = subset[:len(subset)-1]
		}
	}

	dfs(nums, nil, 0, k)
	return res
}
