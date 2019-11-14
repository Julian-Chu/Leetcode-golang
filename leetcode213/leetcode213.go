package leetcode213

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp1 := append([]int{}, nums[:n-1]...)
	dp2 := append([]int{}, nums[1:]...)
	max := 0
	for i := 0; i < n-1; i++ {
		if i >= 2 {
			dp1[i] += dp1[i-2]
			dp2[i] += dp2[i-2]

			if i > 2 {
				if dp1[i-2] < dp1[i-3] {
					dp1[i] += dp1[i-3] - dp1[i-2]
				}
				if dp2[i-2] < dp2[i-3] {
					dp2[i] += dp2[i-3] - dp2[i-2]
				}
			}
		}
		if dp1[i] > max {
			max = dp1[i]
		}
		if dp2[i] > max {
			max = dp2[i]
		}
	}
	return max
}
