package Leetcode_376_Wiggle_Subsequence

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	index := 1

	for index < n && (nums[index]-nums[index-1]) == 0 {
		index++
	}
	if index == n {
		return 1
	}

	gtPrev := true
	if (nums[index] - nums[index-1]) < 0 {
		gtPrev = false
	}
	cnt := 1
	for i := index + 1; i < n; i++ {
		if (gtPrev && nums[i]-nums[i-1] < 0) || (!gtPrev && nums[i]-nums[i-1] > 0) {
			gtPrev = !gtPrev
			cnt++
			continue
		}

	}
	return cnt + 1
}

func wiggleMaxLength_1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	count := 1
	preDiff := 0
	curDiff := 0

	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			preDiff = curDiff
			count++
		}
	}
	return count
}

func wiggleMaxLength_dp(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	n := len(nums)
	dp := [2][]int{}
	// 0-i:max length when peak
	dp[0] = make([]int, n)
	// 0-i:max length when bottom
	dp[1] = make([]int, n)
	dp[0][0] = 1
	dp[1][0] = 1
	for i := 1; i < n; i++ {
		dp[0][i] = 1
		dp[1][i] = 1

		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				dp[1][i] = max(dp[1][i], dp[0][j]+1)
			}

			if nums[j] < nums[i] {
				dp[0][i] = max(dp[0][i], dp[1][j]+1)
			}
		}
	}

	return max(dp[0][n-1], dp[1][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
