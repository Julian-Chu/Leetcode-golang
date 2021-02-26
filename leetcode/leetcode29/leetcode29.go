package leetcode29

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	res, _ := d(int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor))), 1)

	if (divisor > 0 && dividend < 0) || (divisor < 0 && dividend > 0) {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	return res
}

func d(m, n, cnt int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case m >= n && m < n+n:
		return cnt, m - n
	default:
		res, remainder = d(m, n+n, cnt+cnt)
		if remainder >= n {
			return res + cnt, remainder - n
		}

		return
	}
}
