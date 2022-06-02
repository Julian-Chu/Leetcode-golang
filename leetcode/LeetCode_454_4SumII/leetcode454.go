package LeetCode_454_4SumII

func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	sum := make(map[int]int, n*n)

	for i := range A {
		for j := range B {
			sum[A[i]+B[j]]++
		}
	}

	res := 0

	for i := range C {
		for j := range D {
			res += sum[-(C[i] + D[j])]
		}
	}
	return res
}
