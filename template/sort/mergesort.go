package sort

func mergeSort(nums []int, start, end int, tmp []int) {
	if start >= end {
		return
	}
	mid := start + (end-start)>>1
	mergeSort(nums, start, mid, tmp)
	mergeSort(nums, mid+1, end, tmp)
	merge(nums, start, end, tmp)
}

func merge(nums []int, start, end int, tmp []int) {
	mid := start + (end-start)>>1
	l := start
	r := mid + 1
	idx := start

	for l <= mid && r <= end {
		if nums[l] < nums[r] {
			tmp[idx] = nums[l]
			idx++
			l++
		} else {
			tmp[idx] = nums[r]
			idx++
			r++
		}
	}

	for l <= mid {
		tmp[idx] = nums[l]
		idx++
		l++
	}

	for r <= end {
		tmp[idx] = nums[r]
		idx++
		r++
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i]
	}
}
