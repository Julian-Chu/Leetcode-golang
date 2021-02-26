package leetcode378

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]

	for l < r {
		mid := l + (r-l)>>1
		cnt := 0

		for i := range matrix {
			if matrix[i][0] > mid {
				break
			}
			for j := range matrix[i] {
				if matrix[i][j] > mid {
					break
				}
				cnt++
			}
		}

		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
