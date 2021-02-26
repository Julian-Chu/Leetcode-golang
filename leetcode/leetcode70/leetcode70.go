package leetcode70

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	x, y := 1, 1
	for i := 2; i <= n; i++ {
		x, y = x+y, x
	}
	return x
}
