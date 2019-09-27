package leetcode89

import "math"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0, 1}

	for i := 2; i <= n; i++ {
		add := math.Pow(2, float64(i-1))
		cnt := len(res)
		rev := make([]int, cnt)
		for j := cnt - 1; j >= 0; j-- {
			rev[cnt-j-1] = res[j] + int(add)
		}
		res = append(res, rev...)
	}
	return res
}
