package lintcode198

func permutationIndexII(A []int) int64 {
	if len(A) <= 0 {
		return 0
	}

	res := 1
	factorial := 1
	repeat := 1
	counter := make(map[int]int)

	for i := len(A) - 1; i >= 0; i-- {
		if _, ok := counter[A[i]]; !ok {
			counter[A[i]] = 0
		}
		counter[A[i]] += 1
		repeat *= counter[A[i]]
		smaller := 0
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[i] {
				smaller += 1
			}
		}
		res += smaller * factorial / repeat
		factorial *= (len(A) - i)
	}
	return int64(res)
}
