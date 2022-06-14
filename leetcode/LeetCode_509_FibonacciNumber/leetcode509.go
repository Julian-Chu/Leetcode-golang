package LeetCode_509_FibonacciNumber

func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func fib_1(n int) int {
	if n <= 1 {
		return n
	}

	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}
