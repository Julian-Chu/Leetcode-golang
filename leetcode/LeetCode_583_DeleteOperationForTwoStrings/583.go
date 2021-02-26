package leetcode583

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	res := make([][]int, n1+1)
	for i := range res {
		res[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			res[i][j] = max(res[i-1][j], res[i][j-1])
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			}
		}
	}
	return n1 + n2 - res[n1][n2]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
