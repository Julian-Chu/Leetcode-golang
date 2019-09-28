package leetcode90

import (
	"sort"
	"strconv"
	"strings"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	mapper := make(map[string]bool)

	if len(nums) == 0 {
		res = append(res, []int{})
		return res
	}

	var dfs func(int, []int)

	dfs = func(index int, solution []int) {
		if index == len(nums) {
			solutionStr := make([]string, len(solution))
			for i := range solution {
				solutionStr[i] += strconv.Itoa(solution[i])
			}
			key := strings.Join(solutionStr, "-")
			if _, ok := mapper[key]; !ok {
				mapper[key] = true
				res = append(res, solution)
			}
			return
		}
		solution = solution[:len(solution):len(solution)]
		for i := index; i < len(nums); i++ {
			//if i > 0 && nums[i-1] != nums[i] {
			dfs(i+1, solution)
			//}
			dfs(i+1, append(solution, nums[i]))
		}
	}

	dfs(0, []int{})
	return res
}
