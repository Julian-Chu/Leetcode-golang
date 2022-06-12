package LeetCode_134_GasStation

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}

	n := len(gas)

	for i := 0; i < n; i++ {
		cur := 0
		c := 0
		for j := i; j < i+n; j++ {
			if i == j {
				cur += gas[i]
				c = cost[i]
				continue
			}
			idx := j % n
			if cur < c {
				cur = 0
				break
			}
			cur += gas[idx]
			cur -= c
			c = cost[idx]
		}
		if cur >= c {
			return i
		}
	}
	return -1
}

func canCompleteCircuit_greedy(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}

func canCompleteCircuit_2(gas []int, cost []int) int {
	curSum := 0
	min := 1<<31 - 1

	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		if curSum < min {
			min = curSum
		}
	}

	if curSum < 0 {
		return -1
	}

	if min >= 0 {
		return 0
	}

	for i := len(gas) - 1; i >= 0; i-- {
		min += gas[i] - cost[i]
		if min >= 0 {
			return i
		}
	}

	return -1
}
