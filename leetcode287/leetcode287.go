package leetcode287

func findDuplicate(nums []int) int {
	l, r := 1, len(nums)

	for l < r {
		mid := (l + r) / 2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}

		if count > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r

}
