package LeetCode_494_TargetSum

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if (sum+S)%2 == 1 {
		return 0
	}

	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		next := make(map[int]int)
		for sum, cnt := range dp {
			next[sum+nums[i]] += cnt
			next[sum-nums[i]] += cnt
		}
		dp = next
	}
	return dp[S]
}

func findTargetSumWays_1(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if (target+sum)&1 == 1 {
		return 0
	}

	if target > sum || -target > sum {
		return 0
	}

	bagsize := (target + sum) >> 1

	dp := make([]int, bagsize+1)
	dp[0] = 1
	for i := range nums {
		for j := bagsize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagsize]
}
