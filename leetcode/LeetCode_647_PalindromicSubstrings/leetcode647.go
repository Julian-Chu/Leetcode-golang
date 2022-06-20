package LeetCode_647_PalindromicSubstrings

func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += calcPalindrome(s, i, i)
		cnt += calcPalindrome(s, i, i+1)
	}
	return cnt

}

func calcPalindrome(s string, left int, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}

func countSubstrings_dp(s string) int {
	n := len(s)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}
			if j-i <= 1 {
				dp[i][j] = true
				res++
				continue
			}

			if dp[i+1][j-1] {
				res++
				dp[i][j] = true
			}
		}
	}
	return res
}
