package LeetCode_78_Subset

func subsets_1(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	var dfs func(nums, subset []int, start int)
	dfs = func(nums, subset []int, start int) {
		res = append(res, append([]int{}, subset...))
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])
			dfs(nums, subset, i+1)
			subset = subset[:len(subset)-1]
		}
	}

	dfs(nums, make([]int, 0, len(nums)), 0)
	return res
}
