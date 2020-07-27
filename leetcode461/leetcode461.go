package leetcode461

func hammingDistance(x int, y int) int {
	xor := x ^ y
	sum := 0

	for xor > 0 {
		sum += xor & 1
		xor >>= 1
	}

	return sum
}
