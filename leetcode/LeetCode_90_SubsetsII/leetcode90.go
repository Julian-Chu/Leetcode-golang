package LeetCode_90_SubsetsII

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

func subsetsWithDup_1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int
	tmp := make([]int, 0, len(nums))

	var dfs func(nums, tmp []int, start int, visited []bool)
	dfs = func(nums, tmp []int, start int, visited []bool) {
		res = append(res, append([]int{}, tmp...))
		if start == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && visited[i-1] == false {
				continue
			}
			tmp = append(tmp, nums[i])
			visited[i] = true
			dfs(nums, tmp, i+1, visited)
			visited[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(nums, tmp, 0, make([]bool, len(nums)))
	return res
}

func subsetsWithDup_2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int
	tmp := make([]int, 0, len(nums))

	var dfs func(nums, tmp []int, start int)
	dfs = func(nums, tmp []int, start int) {
		res = append(res, append([]int{}, tmp...))
		if start == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] { // be careful i > start not i > 0
				continue
			}
			tmp = append(tmp, nums[i])
			dfs(nums, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(nums, tmp, 0)
	return res
}
