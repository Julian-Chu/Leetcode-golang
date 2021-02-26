package leetcode416

func canPartition(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}
	eqVal := sum >> 1

	dp := make([]bool, eqVal+1)

	dp[0] = true

	for _, num := range nums {
		for i := eqVal; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[eqVal]
}
