package leetcode908

func smallestRangeI(A []int, K int) int {

	// sort.Ints(A)
	// diff := A[len(A)-1] - A[0]
	// if diff > K*2 {
	// 	return diff - K*2
	// }
	// return 0
	if len(A) == 1 {
		return 0
	}
	max := A[0]
	min := A[0]
	for _, e := range A {
		if e > max {
			max = e
		}

		if e < min {
			min = e
		}
	}
	diff := max - min
	if diff > K*2 {
		return diff - K*2
	}
	return 0
}
