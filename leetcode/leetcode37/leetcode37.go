package leetcode37

func solveSudoku(board [][]byte) {
	dfs(board)
}

func dfs(board [][]byte) bool {
	i, j, choices := getLeastChoicesGrid(board)

	if i == -1 {
		return true
	}

	for _, val := range choices {
		board[i][j] = val
		if dfs(board) {
			return true
		}
		board[i][j] = '.'
	}

	return false
}

func getLeastChoicesGrid(board [][]byte) (int, int, []byte) {
	x, y, choices := -1, -1, make([]byte, 10)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			vals := make([]byte, 0, 9)
			for val := 1; val < 10; val++ {
				num := byte(val) + '0'
				if isValid(board, i, j, num) {
					vals = append(vals, num)
				}
			}
			if len(vals) < len(choices) {
				x, y, choices = i, j, vals
			}
		}
	}
	return x, y, choices
}

func isValid(board [][]byte, x, y int, val byte) bool {
	for i := 0; i < 9; i++ {
		if board[x][i] == val {
			return false
		}
		if board[i][y] == val {
			return false
		}
		if board[(x/3)*3+i/3][(y/3)*3+i%3] == val {
			return false
		}
	}
	return true
}
