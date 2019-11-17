package leetcode50

import "math"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 || x == 1 {
		return 1
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}
	res := float64(1)
	n_abs := int(math.Abs(float64(n)))
	for i := 0; i < n_abs; i++ {
		res *= x
	}

	if n < 0 {
		res = 1 / res
	}

	res = math.Floor(res * 100000)
	res /= 100000
	return res
}
