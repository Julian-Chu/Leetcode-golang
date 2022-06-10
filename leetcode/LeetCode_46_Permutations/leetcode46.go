package LeetCode_46_Permutations

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

func permute_1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var res [][]int
	var dfs func(nums, tmp []int, used []bool)
	dfs = func(nums, tmp []int, used []bool) {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			tmp = append(tmp, nums[i])
			used[i] = true
			dfs(nums, tmp, used)
			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(nums, make([]int, 0, len(nums)), make([]bool, len(nums)))

	return res
}
