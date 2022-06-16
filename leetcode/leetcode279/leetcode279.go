package leetcode279

import "math"

func numSquares(n int) int {
	if isSquare(n) {
		return 1
	}
	if n%4 == 0 {
		n /= 4
	}

	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	if n%8 != 7 {
		return 3
	}
	return 4
}

func isSquare(n int) bool {
	c := int(math.Sqrt(float64(n)))

	return c*c == n
}

func numSquares_1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = 1<<31 - 1
	}

	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
