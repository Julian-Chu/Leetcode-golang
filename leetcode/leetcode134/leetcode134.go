package leetcode134

func canCompleteCircuit(gas []int, cost []int) int {
	start, end, sum := len(gas)-1, 0, gas[len(gas)-1]-cost[len(gas)-1]

	for start > end {
		if sum >= 0 {
			sum += gas[end] - cost[end]
			end++
		} else {
			start--
			sum += gas[start] - cost[start]
		}
	}

	if sum >= 0 {
		return start
	}

	return -1
}
