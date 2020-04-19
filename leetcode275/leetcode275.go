package leetcode275

func hIndex(citations []int) int {
	n := len(citations)

	l, r := 0, n-1
	for l < r {
		citations[l], citations[r] = citations[r], citations[l]
		l++
		r--
	}
	for i := 1; i <= n; i++ {
		if citations[i-1] >= i {
			continue
		}

		if citations[i-1] < i {
			return i - 1
		}
	}
	return n
}
