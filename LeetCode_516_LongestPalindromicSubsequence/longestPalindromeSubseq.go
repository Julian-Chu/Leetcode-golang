package leetcode516longestpalindromicsubsequence

func longestPalindromeSubseq(s string) int {
	strLen := len(s)

	subLenMatrix := make([][]int, strLen)
	for i := range subLenMatrix {
		subLenMatrix[i] = make([]int, strLen)
		subLenMatrix[i][i] = 1
	}

	for Len := 2; Len <= strLen; Len++ {
		for i := 0; i+Len-1 < strLen; i++ {
			j := i + Len - 1
			if s[i] == s[j] {
				subLenMatrix[i][j] = subLenMatrix[i+1][j-1] + 2
			} else {
				if subLenMatrix[i][j-1] >= subLenMatrix[i+1][j] {
					subLenMatrix[i][j] = subLenMatrix[i][j-1]
				} else {
					subLenMatrix[i][j] = subLenMatrix[i+1][j]
				}
			}
		}
	}
	return subLenMatrix[0][strLen-1]
}
