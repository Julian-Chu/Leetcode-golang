package leetcode693

func hasAlternatingBits(n int) bool {
	if n == 0 {
		return false
	}
	n = (n >> 1) ^ n

	for n > 0 {
		if n&1 == 0 {
			return false
		}
		n >>= 1
	}
	return true
}
