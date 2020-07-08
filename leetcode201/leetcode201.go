package leetcode201

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	var step uint = 0

	for m != n {
		m >>= 1
		n >>= 1
		step++
	}

	return m << step
}
