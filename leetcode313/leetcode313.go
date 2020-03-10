package leetcode313

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		lower := dp[i-1]
		min := 0
		for j := 1; j < i; j++ {
			for _, p := range primes {
				v := dp[j] * p
				if v > lower {
					if min == 0 {
						min = v
						continue
					}
					if min > v {
						min = v
					}
				}
			}
		}
		dp[i] = min
	}
	return dp[n]
}
