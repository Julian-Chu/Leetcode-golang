package leetcode201

func rangeBitwiseAnd(m int, n int) int {
	v := 0
	var i uint
	for i = 31; i >= 0; i-- {
		digit := 0
		cur := m >> i % 2
		if cur != 0 {
			digit = 1
			for j := m; j <= n; j++ {
				if j>>i%2 == 0 {
					digit = 0
					break
				}
			}

		}
		v *= 2
		v += digit
		if i == 0 {
			break
		}
	}

	return v
}
