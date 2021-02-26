package leetcode1143

func longestCommonSubsequence(text1 string, text2 string) int {
	res := make([][]int, len(text1))

	for i := 0; i < len(text1); i++ {
		res[i] = make([]int, len(text2))
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				res[i][j] = 1
				if i > 0 && j > 0 {
					res[i][j] = max(1, res[i-1][j-1]+1)
				}
			} else {
				if i == 0 && j > 0 {
					res[i][j] = res[i][j-1]
				} else if i > 0 && j == 0 {
					res[i][j] = res[i-1][j]
				} else if i > 0 && j > 0 {
					res[i][j] = max(res[i][j-1], res[i-1][j])
				}
			}
		}
	}
	maxLen := 0
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if res[i][j] > maxLen {
				maxLen = res[i][j]
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
