package leetcode29

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}
	cnt := 0
	abs_dividend := int(math.Abs(float64(dividend)))
	d := int(math.Abs(float64(divisor)))
	for abs_dividend >= d {
		abs_dividend -= d
		cnt++
	}

	if (divisor > 0 && dividend > 0) || (divisor < 0 && dividend < 0) {
		if cnt > math.MaxInt32 {
			return math.MaxInt32
		}
		return cnt
	} else {
		if cnt < math.MinInt32 {
			return math.MaxInt32
		}
		return -cnt
	}
}
