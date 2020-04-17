package leetcode494

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 0
	var dfs func(int, int)
	dfs = func(idx, sum int) {
		plus := sum + nums[idx]
		minus := sum - nums[idx]
		if idx == n-1 {
			if plus == S {
				res++
			}
			if minus == S {
				res++
			}
			return
		}

		dfs(idx+1, plus)
		dfs(idx+1, minus)
	}
	dfs(0, 0)
	return res
}
