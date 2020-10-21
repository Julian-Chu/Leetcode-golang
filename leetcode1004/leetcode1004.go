package leetcode1004

func longestOnes(A []int, K int) int {
	start := 0
	zeros := 0
	res := 0

	for i := range A {
		if A[i] == 0 {
			zeros++
			if zeros > K {
				zeros--
				for A[start] != 0 {
					start++
				}
				start++
			}
		}
		res = max(res, i-start+1)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
