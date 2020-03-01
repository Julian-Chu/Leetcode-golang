package leetcode376

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp_cnt := make([][]int, n)
	dp_diff := make([][]int, n)
	dp_last_idx := make([][]int, n)

	for i := range dp_cnt {
		dp_cnt[i] = make([]int, n)
		dp_diff[i] = make([]int, n)
		dp_last_idx[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if j == i+1 {
				if nums[i] == nums[j] {
					dp_cnt[i][j] = 0
					dp_diff[i][j] = 0
				} else {
					dp_cnt[i][j] = 1
					dp_diff[i][j] = nums[j] - nums[i]
				}
				dp_last_idx[i][j] = j
				continue
			}

			prevNum := nums[dp_last_idx[i][j-1]]
			diff := nums[j] - prevNum
			if diff == 0 {
				dp_diff[i][j] = dp_diff[i][j-1]
				dp_last_idx[i][j] = j
				dp_cnt[i][j] = dp_cnt[i][j-1]
				continue
			}
			if dp_diff[i][j-1] == 0 {
				if diff != 0 {
					dp_cnt[i][j] = 1
				}
				dp_diff[i][j] = diff
				dp_last_idx[i][j] = j
				continue
			}

			if dp_diff[i][j-1] > 0 {
				if diff > 0 {
					if nums[j] >= prevNum {
						dp_last_idx[i][j] = j
						dp_diff[i][j] = diff
					} else {
						dp_last_idx[i][j] = j - 1
						dp_diff[i][j] = dp_diff[i][j-1]
					}
					dp_cnt[i][j] = dp_cnt[i][j-1]
				} else {
					dp_last_idx[i][j] = j
					dp_diff[i][j] = diff
					dp_cnt[i][j] = dp_cnt[i][j-1] + 1
				}
				continue
			}

			if dp_diff[i][j-1] < 0 {
				if diff < 0 {
					if nums[j] <= prevNum {
						dp_last_idx[i][j] = j
						dp_diff[i][j] = diff
					} else {
						dp_last_idx[i][j] = j - 1
						dp_diff[i][j] = dp_diff[i][j-1]
					}
					dp_cnt[i][j] = dp_cnt[i][j-1]
				} else {
					dp_last_idx[i][j] = j
					dp_diff[i][j] = diff
					dp_cnt[i][j] = dp_cnt[i][j-1] + 1
				}
				continue
			}

		}
	}
	return dp_cnt[0][n-1] + 1
}
