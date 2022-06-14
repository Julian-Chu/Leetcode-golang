package LeetCode_746_MinCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(dp)-1]

}

func minCostClimbingStairs_1(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[len(dp)-1], dp[len(dp)-2])

}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
