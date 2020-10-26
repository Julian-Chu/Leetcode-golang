package leetcode1423

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	notTaken := n - k
	sum := 0
	minDis := 0
	dis := 0
	for i := 0; i < n; i++ {
		sum += cardPoints[i]
		dis += cardPoints[i]
		if i-notTaken < 0 {
			minDis = dis
			continue
		}
		dis -= cardPoints[i-notTaken]
		minDis = min(dis, minDis)
	}

	minDis = min(dis, minDis)
	return sum - minDis
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
