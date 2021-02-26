package lintcode152

/**
 * @param n: Given the range of numbers
 * @param k: Given the numbers of combinations
 * @return: All the combinations of k numbers out of 1..n
 */
func combine(n int, k int) [][]int {
	var result [][]int
	var comb []int
	dfs(n, 1, &comb, &result, k)
	return result
}

func dfs(n, start int, comb *[]int, result *[][]int, k int) {
	if k == 0 {
		clone := append((*comb)[:0:0], *comb...)
		*result = append(*result, clone)
		return
	}

	for i := start; i < n+1; i++ {
		*comb = append(*comb, i)
		dfs(n, i+1, comb, result, k-1)
		*comb = (*comb)[:len(*comb)-1]
	}
}
