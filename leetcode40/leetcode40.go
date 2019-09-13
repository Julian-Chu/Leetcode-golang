package leetcode40

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0, 100)
	dfs(candidates, target, &res, []int{})

	return res
}

func dfs(candidates []int, target int, res *[][]int, solution []int) {
	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	if target == candidates[0] {
		solution = append(solution, candidates[0])
		*res = append(*res, solution)
		return
	}

	solution = solution[:len(solution):len(solution)]
	dfs(next(candidates), target, res, solution)
	solution = append(solution, candidates[0])
	dfs(candidates[1:], target-candidates[0], res, solution)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}

//func combinationSum2(candidates []int, target int) [][]int {
//	sort.Ints(candidates)
//	res := make([][]int, 0, 100)
//	mapper := make(map[string]bool)
//	dfs(candidates, target, &res, []int{}, mapper)
//
//	return res
//}
//
//func dfs(candidates []int, target int, res *[][]int, solution []int, mapper map[string]bool) {
//
//	if len(candidates) == 0 || target < candidates[0] {
//		return
//	}
//
//	if target == candidates[0] {
//		solution = append(solution, candidates[0])
//		str := make([]string, len(solution))
//		for i := range solution {
//			txt := strconv.Itoa(solution[i])
//			str[i] = txt
//		}
//
//		key := strings.Join(str, ",")
//		if _, ok := mapper[key]; !ok {
//			mapper[key] = true
//			*res = append(*res, solution)
//		}
//
//		return
//	}
//
//	solution = solution[:len(solution):len(solution)]
//	dfs(candidates[1:], target, res, solution, mapper)
//	solution = append(solution, candidates[0])
//	dfs(candidates[1:], target-candidates[0], res, solution, mapper)
//}
