package leetcode91

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0]-'0' == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	lastDigit, lastTwoDigits := s[0]-'0', uint8(0)

	for i := 1; i < n; i++ {
		lastDigit, lastTwoDigits = s[i]-'0', lastDigit*10+s[i]-'0'
		if lastDigit > 0 {
			dp[i] += dp[i-1]
		}
		if lastTwoDigits >= 10 && lastTwoDigits <= 26 {
			if i == 1 {
				dp[i]++
			} else {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n-1]

}
