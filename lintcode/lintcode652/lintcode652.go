package lintcode652

/**
 * @param n: An integer
 * @return: a list of combination
 */
import (
	"math"
	"sort"
)

func getFactors(n int) [][]int {
	result := [][]int{}
	res := []int{}
	dfs(n, 2, &res, &result)
	sort.Slice(result, func(i, j int) bool {
		minLen := len(result[i])
		if len(result[j]) < len(result[i]) {
			minLen = len(result[j])
		}

		for k := 0; k < minLen; k++ {
			if result[i][k] < result[j][k] {
				return true
			}

			if result[i][k] > result[j][k] {
				return false
			}
		}

		return false
	})
	return result
}

func dfs(n, start int, res *[]int, result *[][]int) {
	if len(*res) > 0 {
		clone := *res
		clone = append(clone[:0:0], clone...)
		clone = append(clone, n)
		*result = append(*result, clone)
	}

	up := int(math.Sqrt(float64(n)))
	for i := start; i < up+1; i++ {
		if n%i == 0 {
			*res = append(*res, i)
			dfs(n/i, i, res, result)
			*res = (*res)[:len(*res)-1]
		}
	}
}
