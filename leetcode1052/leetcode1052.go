package leetcode1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum, max := 0, 0

	for i := range customers {
		sum += customers[i] * (1 - grumpy[i])

	}

	n := len(grumpy)
	i, j, satisfied := 0, 0, 0
	for j-i < X {
		satisfied += customers[j] * grumpy[j]
		j++
	}
	max = satisfied
	for j < n {
		satisfied -= customers[i] * grumpy[i]
		satisfied += customers[j] * grumpy[j]
		if satisfied > max {
			max = satisfied
		}
		i++
		j++
	}

	return sum + max
}
