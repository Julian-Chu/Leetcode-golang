package leetcode240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])

	k := m
	if m > n {
		k = n
	}
	start := k - 1
	for i := 0; i < k; i++ {
		if matrix[i][i] == target {
			return true
		}
		if target < matrix[i][i] {
			if i == 0 {
				return false
			}
			start = i - 1
			break
		}
	}

	for idx := 0; idx < start+1; idx++ {
		for i := idx; i < n; i++ {
			if matrix[idx][i] == target {
				return true
			}
		}

		if idx+1 < m {
			for i := 0; i < idx; i++ {
				if matrix[idx+1][i] == target {
					return true
				}
			}
		}
		for j := idx; j < m; j++ {
			if matrix[j][idx] == target {
				return true
			}
		}

		if idx+1 < n {
			for j := 0; j < idx; j++ {
				if matrix[j][idx+1] == target {
					return true
				}
			}
		}
	}

	return false
}
