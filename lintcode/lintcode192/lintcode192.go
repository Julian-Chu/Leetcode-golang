package lintcode192

/**
 * @param s: A string
 * @param p: A string includes "?" and "*"
 * @return: is Match?
 */
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 1; i < m+1; i++ {
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}

		}
	}

	return dp[n][m]

}

/// dfs + memo
//import (
//"strconv"
//)
//func isMatch (s string, p string) bool {
//	visited := make(map[string]bool)
//	return dfs(s, p, 0, 0, visited)
//}
//
//func dfs(s string, p string, s_index, p_index int, visited map[string]bool) bool{
//	// idx1 := strconv.Itoa(s_index)
//	// idx2 := strconv.Itoa
//	key := strconv.Itoa(s_index) + "-" + strconv.Itoa(p_index)
//	if v, ok:= visited[key];ok{
//		return v
//	}
//	if p_index >= len(p){
//		return s_index >= len(s)
//	}
//
//	if s_index >=len(s){
//		// return p[p_index:] == "*"
//		for j:=p_index; j <len(p);j++{
//			if p[j] != '*'{
//				return false
//			}
//		}
//		return true
//	}
//
//	if s[s_index] == p[p_index] || p[p_index] == '?'{
//		visited[key] = dfs(s, p, s_index +1, p_index+1, visited)
//		return visited[key]
//	}
//
//	if p[p_index] == '*'{
//		for i:=s_index; i <len(s)+1;i++{
//			if dfs(s, p, i, p_index+1, visited){
//				visited[key] = true
//				return visited[key]
//			}
//		}
//	}
//
//	return false
//}
