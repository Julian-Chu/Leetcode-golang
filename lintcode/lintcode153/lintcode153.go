package lintcode153

/**
 * @param num: Given the candidate numbers
 * @param target: Given the target number
 * @return: All the combinations that sum to target
 */
import "sort"

func combinationSum2(num []int, target int) [][]int {
	result := [][]int{}
	comb := []int{}
	sort.Ints(num)
	dfs(num, 0, &comb, &result, target)
	return result
}

func dfs(num []int, start int, comb *[]int, result *[][]int, target int) {
	if target == 0 {
		newComb := make([]int, len(*comb))
		copy(newComb, *comb)
		*result = append(*result, newComb)
		return
	}

	if start == len(num) {
		return
	}

	for i := start; i < len(num); i++ {
		if num[i] > target {
			break
		}

		if i != start && num[i-1] == num[i] {
			continue
		}

		*comb = append(*comb, num[i])
		dfs(num, i+1, comb, result, target-num[i])
		*comb = (*comb)[:len(*comb)-1]
	}
}
