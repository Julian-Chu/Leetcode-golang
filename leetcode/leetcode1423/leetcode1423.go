package leetcode1423

func maxScore(cardPoints []int, k int) int {
	sum := 0
	n := len(cardPoints)

	for i := 1; i <= k; i++ {
		sum += cardPoints[n-i]
	}

	max := sum
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
		sum -= cardPoints[n-k+i]
		if sum > max {
			max = sum
		}
	}
	return max
}
