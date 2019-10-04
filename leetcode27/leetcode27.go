package leetcode27

func removeElement(nums []int, val int) int {
	noval := true
	for _, v := range nums {
		if v == val {
			noval = false
		}
	}

	if noval {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] != val {
			l++
		} else {
			if nums[r] == val {
				r--
			} else {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
	}
	return l
}
