package leetcode120

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			min := triangle[i+1][j]
			if triangle[i+1][j] > triangle[i+1][j+1] {
				min = triangle[i+1][j+1]
			}
			triangle[i][j] += min
		}
	}

	return triangle[0][0]
}
