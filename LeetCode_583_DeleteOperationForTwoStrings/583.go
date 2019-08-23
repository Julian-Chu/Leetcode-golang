package leetcode583

import (
	"strings"
)

func minDistance(word1 string, word2 string) int {
	if strings.Compare(word1, word2) == 0 {
		return 0
	}
	res := make([][]int, len(word1))
	maxLen := 0
	for i := range res {
		res[i] = make([]int, len(word2))
		for j := range res[i] {
			if word1[i] == word2[j] {
				res[i][j] = 1
				if i > 0 && j > 0 {
					res[i][j] = max(1, res[i-1][j-1]+1)
				}
				if maxLen < res[i][j] {
					maxLen = res[i][j]
				}
				continue
			}
			if i > 0 && j > 0 {
				res[i][j] = max(res[i-1][j], res[i][j-1])
			} else if i > 0 && j == 0 {
				res[i][j] = res[i-1][j]
			} else if i == 0 && j > 0 {
				res[i][j] = res[i][j-1]
			}
			if maxLen < res[i][j] {
				maxLen = res[i][j]
			}
		}
	}
	return len(word1) + len(word2) - maxLen*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
