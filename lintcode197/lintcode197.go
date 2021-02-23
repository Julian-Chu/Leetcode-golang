package lintcode197

func permutationIndex(A []int) int64 {
	if len(A) == 0 {
		return 0
	}

	factorial := 1
	res := 1
	for i := len(A) - 1; i >= 0; i-- {
		smaller := 0
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				smaller += 1
			}
		}
		res += smaller * factorial
		factorial *= (len(A) - i)
	}
	return int64(res)
}
