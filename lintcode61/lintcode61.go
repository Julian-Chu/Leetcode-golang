package lintcode61

/**
 * @param A: an integer sorted array
 * @param target: an integer to be inserted
 * @return: a list of length 2, [index1, index2]
 */
func searchRange(A []int, target int) []int {
	if A == nil || len(A) == 0 {
		return []int{-1, -1}
	}

	start, end := 0, len(A)-1
	first, last := -1, -1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] >= target {
			end = mid
		} else {
			start = mid
		}
	}

	if A[start] == target {
		first = start
	} else if A[end] == target {
		first = end
	}

	if first == -1 {
		return []int{-1, -1}
	}

	start, end = first, len(A)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}

	last = start
	if A[end] == target {
		last = end
	}

	return []int{first, last}
}
