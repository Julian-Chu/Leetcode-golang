package leetcode454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	m1 := make(map[int]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m1[A[i]+B[j]]++
		}
	}
	m2 := make(map[int]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m2[C[i]+D[j]]++
		}
	}

	res := 0
	for sum1, cnt1 := range m1 {
		sum2 := 0 - sum1
		if cnt2, ok := m2[sum2]; ok {
			res += cnt1 * cnt2
		}
	}

	return res
}
