package lintcode183

/**
 * @param L: Given n pieces of wood with length L[i]
 * @param k: An integer
 * @return: The maximum length of the small pieces
 */
// import "fmt"
func woodCut(L []int, k int) int {
	if L == nil || len(L) == 0 {
		return 0
	}
	start, end := 1, maxCutLen(L, k)
	if start > end {
		return 0
	}
	for start+1 < end {
		cutLen := start + (end-start)/2
		pcs := cutToPieces(L, cutLen)
		if pcs < k {
			end = cutLen
		} else {
			start = cutLen
		}
	}

	if cutToPieces(L, end) == k {
		return end
	}
	return start
}

func cutToPieces(L []int, cutLen int) int {
	sum := 0
	for _, l := range L {
		sum += l / cutLen
	}
	return sum
}

func maxCutLen(L []int, k int) int {
	maxLen := L[0]
	sum := 0
	for _, l := range L {
		sum += l
		if l > maxLen {
			maxLen = l
		}
	}

	if maxLen < sum/k {
		return maxLen
	}
	return sum / k
}
