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
