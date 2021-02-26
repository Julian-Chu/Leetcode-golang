package lintcode62

/**
 * @param A: an integer rotated sorted array
 * @param target: an integer to be searched
 * @return: an integer
 */
func search(A []int, target int) int {
	if A == nil || len(A) == 0 {
		return -1
	}

	lastNum := A[len(A)-1]
	start, end := 0, len(A)-1

	for start+1 < end {
		mid := start + (end-start)/2
		if target <= lastNum { // target in right
			if A[mid] > lastNum {
				start = mid
				continue
			}

			// mid in right
			if A[mid] == target {
				return mid
			} else if A[mid] < target {
				start = mid
			} else {
				end = mid
			}

		} else { //target in left
			if A[mid] <= lastNum { // mid in right
				end = mid
				continue
			}

			// mid in left
			switch {
			case A[mid] == target:
				return mid
			case A[mid] < target:
				start = mid
			default:
				end = mid
			}
		}
	}
	if A[start] == target {
		return start
	}
	if A[end] == target {
		return end
	}
	return -1
}
