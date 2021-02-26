package leetcode46

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	dfs(nums, make([]int, 0), &res)
	return res
}

func dfs(nums, solution []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, solution)
	}

	for i := range nums {
		solution = solution[:len(solution):len(solution)]
		newNums := append([]int{}, nums[0:i]...)
		newNums = append(newNums, nums[i+1:]...)
		if i == 0 {
			newNums = append([]int{}, nums[1:]...)
		}
		dfs(newNums, append(solution, nums[i]), res)
	}
}
