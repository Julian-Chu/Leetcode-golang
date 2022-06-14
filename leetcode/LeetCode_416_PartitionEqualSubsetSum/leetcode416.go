package LeetCode_416_PartitionEqualSubsetSum

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

func canPartition_dp(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = 0
	}

	for j := nums[0]; j < target+1; j++ {
		dp[0][j] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < target+1; j++ {
			if j >= nums[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}

	return dp[len(nums)-1][target] == target
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func canPartition_dp_1d(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}

	return dp[target] == target
}

func canPartition_memo(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	target := sum >> 1
	n := len(nums)
	memo := map[[2]int]bool{}
	var dfs func(nums []int, idx, subsetSum int) bool
	dfs = func(nums []int, idx, subsetSum int) bool {
		if subsetSum == 0 {
			return true
		}
		if idx == n || subsetSum < 0 {
			return false
		}

		memoKey := [2]int{idx, subsetSum}
		if _, ok := memo[memoKey]; ok {
			return memo[memoKey]
		}

		res := dfs(nums, idx+1, subsetSum) || dfs(nums, idx+1, subsetSum-nums[idx])

		memo[memoKey] = res
		return res
	}
	return dfs(nums, 0, target)
}
