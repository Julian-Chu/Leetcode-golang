package leetcode1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	satisfied, dis, maxDis := 0, 0, 0

	for i := 0; i < n; i++ {
		satisfied += customers[i] * (1 - grumpy[i])
		dis += customers[i] * grumpy[i]
		if i-X >= 0 {
			dis -= customers[i-X] * grumpy[i-X]
		}
		maxDis = max(maxDis, dis)
	}

	return satisfied + maxDis
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
