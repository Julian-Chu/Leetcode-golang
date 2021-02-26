package leetcode396

func maxRotateFunction(A []int) int {
	//F(i)=F(i-1) + sum - n x A[n-i]
	F, sum := 0, 0
	n := len(A)
	for i := 0; i < n; i++ {
		F += i * A[i]
		sum += A[i]
	}

	max := F

	for i := 1; i < n; i++ {
		F = F + sum - n*A[n-i]
		if F > max {
			max = F
		}
	}

	return max
}
