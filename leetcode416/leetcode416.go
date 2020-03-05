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
	dp_slice := make([][]int, eqVal+1)

	for i := range dp {
		if i == 0 {
			dp_slice[i] = append([]int{}, nums...)
			continue
		}

	loop:
		for j := i - 1; j >= 0; j-- {
			v := i - j
			for idx := range dp_slice[j] {
				if dp[j] == false {
					continue
				}
				if dp_slice[j][idx] == v {
					lastIdx := len(dp_slice[j]) - 1
					dp_slice[j][lastIdx], dp_slice[j][idx] = dp_slice[j][idx], dp_slice[j][lastIdx]
					dp_slice[i] = make([]int, lastIdx)
					copy(dp_slice[i], dp_slice[j][:lastIdx])
					dp[i] = true
					break loop
				}
			}
		}
		if dp[i] == false {
			dp_slice[i] = append([]int{}, nums...)
		}
	}

	return dp[eqVal]
}
