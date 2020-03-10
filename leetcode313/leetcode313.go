package leetcode313

func nthSuperUglyNumber(n int, primes []int) int {
	if len(primes) == 0 {
		return 1
	}
	max := 1
	for i := 0; i < n-1; i++ {
		max *= primes[0]
	}
	dp := make([]bool, max+1)
	dp[1] = true
	cnt := 1
	i := 2
	for cnt < n {
		for idx := range primes {
			if i%primes[idx] == 0 && dp[i/primes[idx]] == true {
				dp[i] = true
				cnt++
				if cnt == n {
					return i
				}
				break
			}
		}
		i++
	}

	return 1
}
