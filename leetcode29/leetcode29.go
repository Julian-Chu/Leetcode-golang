package leetcode29

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}

	res := divInt64(int64(math.Abs(float64(dividend))), int64(math.Abs(float64(divisor))))

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		res = -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if -res < math.MinInt32 {
		return math.MinInt32
	}
	return int(res)
}

func divInt64(dividend int64, divisor int64) int64 {
	if dividend < divisor {
		return 0
	}
	cnt := 1
	temp := divisor
	for dividend > (temp << 1) {
		temp = temp << 1
		cnt = cnt << 1
	}

	return int64(cnt) + divInt64(dividend-temp, divisor)
}
