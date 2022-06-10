package LeetCode_47_PermutationsII

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
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

			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
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

//bad
//func permuteUnique(nums []int) [][]int {
//	if len(nums) == 0 {
//		return [][]int{}
//	}
//	sort.Ints(nums)
//	res := make([][]int, 0)
//
//	rec(nums, []int{}, &res)
//	return res
//}
//
//func rec(nums, solution []int, res *[][]int) {
//	if len(nums) == 0 {
//		*res = append(*res, solution)
//		return
//	}
//
//	mapper := make(map[int]bool)
//
//	for i := range nums {
//		//if i > 0 && nums[i-1] == nums[i] {
//		if v := mapper[nums[i]]; v {
//			continue
//		}
//		mapper[nums[i]] = true
//		newNums := append([]int{}, nums[1:]...)
//		if i != 0 {
//			newNums = append(append([]int{}, nums[0:i]...), nums[i+1:]...)
//		}
//		solution = solution[:len(solution):len(solution)]
//		rec(newNums, append(solution, nums[i]), res)
//	}
//}
