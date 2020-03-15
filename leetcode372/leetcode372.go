package leetcode372

func superPow(a int, b []int) int {
	cnt := 0
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		cyc := 1
		for j := 0; j < cnt; j++ {
			cyc *= 10
		}

		for k := 0; k < cyc*b[i]; k++ {
			res *= a
			res %= 1337
		}
		cnt++
	}

	return res
}
