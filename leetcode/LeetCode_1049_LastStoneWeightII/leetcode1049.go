package LeetCode_1049_LastStoneWeightII

func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := range stones {
		sum += stones[i]
	}

	target := sum / 2

	dp := make([]int, target+1)

	for i := range stones {
		for j := target; j >= 0; j-- {
			if j >= stones[i] {
				dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			}
		}
	}

	return sum - dp[target]*2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
