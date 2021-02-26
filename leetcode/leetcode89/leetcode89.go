package leetcode89

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0, 1}
	base := 1
	for i := 2; i <= n; i++ {
		base *= 2
		cnt := len(res)
		rev := make([]int, cnt)
		for j := cnt - 1; j >= 0; j-- {
			rev[cnt-j-1] = res[j] + base
		}
		res = append(res, rev...)
	}
	return res
}
