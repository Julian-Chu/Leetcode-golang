package leetcode134

func canCompleteCircuit(gas []int, cost []int) int {
	end := len(gas) - 1
	for start := range gas {
		i := start
		sum := gas[i] - cost[i]
		i++
		if i > end {
			i = 0
		}
		for sum >= 0 {
			sum += gas[i] - cost[i]
			if sum < 0 {
				break
			}

			if i == start {
				return start
			}
			i++
			if i > end {
				i = 0
			}
		}
	}

	return -1
}
