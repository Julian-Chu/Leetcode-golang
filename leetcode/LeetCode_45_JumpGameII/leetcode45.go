package LeetCode_45_JumpGameII

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	curDist := 0
	nextDist := 0
	ans := 0
	for i := 0; i < n; i++ {
		nextDist = max(nums[i]+i, nextDist)
		if i == curDist {
			if curDist == n-1 {
				break
			}
			ans++
			curDist = nextDist
			if nextDist >= n-1 {
				break
			}
		}
	}
	return ans
}

func jump_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	curDist := 0
	nextDist := 0
	ans := 0
	// stop at n-2
	for i := 0; i < n-1; i++ {
		nextDist = max(nums[i]+i, nextDist)
		if i == curDist {
			curDist = nextDist
			ans++
		}
	}
	return ans
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func jump_dp(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
