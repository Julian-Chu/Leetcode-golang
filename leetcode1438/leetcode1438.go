package leetcode1438

func longestSubarray(nums []int, limit int) int {
	n := len(nums)
	maxd := make([]int, 0, n)
	mind := make([]int, 0, n)
	i := 0

	for _, num := range nums {
		for len(maxd) > 0 && num > maxd[len(maxd)-1] {
			maxd = maxd[:len(maxd)-1]
		}
		for len(mind) > 0 && num < mind[len(mind)-1] {
			mind = mind[:len(mind)-1]
		}

		maxd = append(maxd, num)
		mind = append(mind, num)

		if maxd[0]-mind[0] > limit {
			if maxd[0] == nums[i] {
				copy(maxd, maxd[1:])
				maxd = maxd[:len(maxd)-1]
			}

			if mind[0] == nums[i] {
				copy(mind, mind[1:])
				mind = mind[:len(mind)-1]
			}

			i++
		}
	}
	return n - i
}
