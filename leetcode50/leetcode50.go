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

	res := pow(x, n)

	if n < 0 {
		res = 1 / res
	}

	res = math.Floor(res * 100000)
	res /= 100000
	return res
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		return pow(x*x, n/2)
	} else {
		return x * pow(x*x, n/2)
	}
}
