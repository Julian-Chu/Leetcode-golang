package leetcode372

func superPow(a int, b []int) int {
	n := len(b)
	base := 1337

	if n == 0 {
		return 1
	}
	// a^k mode base
	powmod := func(a, k int) int {
		a %= base
		res := 1

		for i := 0; i < k; i++ {
			res = (res * a) % base
		}
		return res
	}
	return (powmod(superPow(a, b[:n-1]), 10) * powmod(a, b[n-1])) % base
}
