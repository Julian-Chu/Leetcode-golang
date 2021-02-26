package leetcode96

func numTrees(n int) int {
	if n <= 2 {
		return n
	}

	res := make([]int, n+1)
	res[0], res[1] = 1, 1

	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			res[i] += res[j-1] * res[i-j]
		}
	}
	return res[n]
}
