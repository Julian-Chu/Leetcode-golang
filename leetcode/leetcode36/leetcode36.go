package leetcode36

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			if val > '9' || val < '1' {
				return false
			}

			if !m[val] {
				m[val] = true
				continue
			}
			return false
		}
	}

	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			val := board[j][i]
			if val == '.' {
				continue
			}
			if val > '9' || val < '1' {
				return false
			}
			if !m[val] {
				m[val] = true
				continue
			}
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			mapper := make(map[byte]bool)
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					val := board[m+i*3][n+j*3]
					if val == '.' {
						continue
					}
					if val > '9' || val < '1' {
						return false
					}
					if !mapper[val] {
						mapper[val] = true
						continue
					}
					return false
				}
			}
		}
	}
	return true
}
