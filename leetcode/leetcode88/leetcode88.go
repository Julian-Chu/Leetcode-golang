package leetcode88

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums3 := append([]int{}, nums1[:m]...)
	j, k := 0, 0
	for i := 0; i < m+n; i++ {
		switch {
		case j >= m:
			nums1[i] = nums2[k]
			k++
		case k >= n:
			nums1[i] = nums3[j]
			j++
		default:
			if nums2[k] >= nums3[j] {
				nums1[i] = nums3[j]
				j++
			} else {
				nums1[i] = nums2[k]
				k++
			}
		}
	}
}

func merge_1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)

	quicksort(nums1, 0, len(nums1)-1)
}

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

func merge_2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)

	tmp := make([]int, m+n)
	mergeSort(nums1, 0, len(nums1)-1, tmp)
}

func mergeSort(nums []int, start, end int, tmp []int) {
	if start >= end {
		return
	}
	mid := start + (end-start)>>1
	mergeSort(nums, start, mid, tmp)
	mergeSort(nums, mid+1, end, tmp)
	mergeSorted(nums, start, end, tmp)
}

func mergeSorted(nums []int, start, end int, tmp []int) {
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
