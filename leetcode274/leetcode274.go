package leetcode274

func hIndex(citations []int) int {
	qsort(citations)

	for i, num := range citations {
		if i >= num {
			return i
		}
	}
	return len(citations)
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	l := 0
	r := len(nums) - 1
	mid := nums[0]

	for i := 1; i <= r; {
		if nums[i] < mid {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		}
	}
	qsort(nums[:r])
	qsort(nums[r+1:])
}
