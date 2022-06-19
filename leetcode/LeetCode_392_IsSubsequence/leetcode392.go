package LeetCode_392_IsSubsequence

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	cnt := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			cnt++
		}
		j++
	}

	return cnt == len(s)
}

func isSubsequence_dp(s string, t string) bool {
	if s == "" {
		return true
	}

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}
