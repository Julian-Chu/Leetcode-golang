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

//func longestPalindromeSubseq (s string) int {
//	if s == ""{
//		return 0
//	}
//
//	n := len(s)
//	dp := make([][]int, n)
//	for i := range dp{
//		dp[i] = make([]int,n)
//		dp[i][i] = 1
//	}
//
//	for l:= 2; l < n+1; l++{
//		for i:=0; i < n-l+1 ; i++{
//			j := i+l-1
//			if s[i]==s[j]{
//				dp[i][j] = dp[i+1][j-1] + 2
//			}else{
//				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
//			}
//		}
//
//	}
//	return dp[0][n-1]
//}
//
//func max(a,b int)int{
//	if a>b{
//		return a
//	}
//	return b
//}
