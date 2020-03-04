package leetcode413

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	dp := make([]int, n)

	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i]++
			dp[i] += dp[i-1]
		}
	}
	sum := 0
	for i := range dp {
		sum += dp[i]
	}

	return sum
}
