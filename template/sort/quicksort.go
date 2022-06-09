package sort

func quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}

	l, r := start, end
	pivot := nums[(start + (end-start)>>1)]

	for l <= r {
		for l <= r && nums[l] < pivot {
			l++
		}

		for l <= r && nums[r] > pivot {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	quicksort(nums, start, r)
	quicksort(nums, l, end)
}
