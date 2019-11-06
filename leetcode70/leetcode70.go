package leetcode70

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
