package LeetCode_55_JumpGame

func canJump(nums []int) bool {
	n := len(nums) - 1
	if n == 0 {
		return true
	}
	var prevJumpPoint func(int) bool

	prevJumpPoint = func(cur int) bool {
		for i := cur - 1; i >= 0; i-- {
			if cur-i <= nums[i] {
				if i == 0 {
					return true
				}
				return prevJumpPoint(i)
			}
		}
		return false
	}
	return prevJumpPoint(n)
}

func canJump_greedy(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	cover := nums[0]
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func canJump_dp(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	n := len(nums)
	dp := make([]bool, n)

	dp[0] = true

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}

	return dp[n-1]
}
