package leetcode264

func nthUglyNumber(n int) int {
	res := []int{1}
	i2, i3, i5 := 0, 0, 0

	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for len(res) < n {
		m2, m3, m5 := res[i2]*2, res[i3]*3, res[i5]*5

		m := min(min(m2, m3), m5)

		if m == m2 {
			i2++
		}

		if m == m3 {
			i3++
		}

		if m == m5 {
			i5++
		}
		res = append(res, m)
	}

	return res[n-1]
}
