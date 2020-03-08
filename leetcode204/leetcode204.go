package leetcode204

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]bool, n)

	for i := range dp {
		dp[i] = true
	}
	dp[0] = false
	dp[1] = false

	for i := 2; i < n; i++ {
		if dp[i] == false {
			continue
		}

		for j := i * i; j < n; {
			dp[j] = false
			j += i
		}
	}

	cnt := 0

	for _, v := range dp {
		if v {
			cnt++
		}
	}

	return cnt
}
