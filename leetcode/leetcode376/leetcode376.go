package leetcode376

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
