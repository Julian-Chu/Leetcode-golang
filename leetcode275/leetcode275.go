package leetcode275

func hIndex(citations []int) int {
	return bSearchIndex(citations, 0, len(citations)-1)
}

func bSearchIndex(c []int, l int, r int) int {
	for l < r {
		mid := l + (r-l)/2
		if len(c)-mid > c[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if len(c) > 0 && len(c)-l > c[l] {
		return 0
	}
	return len(c) - l
}
