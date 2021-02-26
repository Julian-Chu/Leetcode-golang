package lintcode77

func longestCommonSubsequence(A string, B string) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	n, m := len(A), len(B)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if A[i-1] == B[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func longestCommonSubsequence (A string, B string) int {
//	if len(A)==0 || len(B)==0{
//		return 0
//	}
//
//	n, m := len(A), len(B)
//
//	dp:=make([][]int, 2)
//	for i :=range(dp){
//		dp[i] = make([]int, m+1)
//	}
//
//	for i:=1; i<n+1; i++{
//		for j:=1; j<m+1; j++{
//			dp[i%2][j] = max(dp[(i-1)%2][j], dp[i%2][j-1])
//			if A[i-1] == B[j-1]{
//				dp[i%2][j] = max(dp[i%2][j], dp[(i-1)%2][j-1]+1)
//			}
//		}
//	}
//
//	return dp[n%2][m]
//}
