package LeetCode_51_N_Queens

func solveNQueens(n int) [][]string {
	result := [][]string{}
	cols := []int{}
	dfs(n, &cols, &result)
	return result
}

func dfs(n int, cols *[]int, result *[][]string) {
	row := len(*cols)
	if row == n {
		*result = append(*result, drawChessBoard(*cols))
		return
	}

	for col := 0; col < n; col++ {
		if !isValid(*cols, row, col) {
			continue
		}
		*cols = append(*cols, col)
		dfs(n, cols, result)
		*cols = (*cols)[:len(*cols)-1]
	}
}

func isValid(cols []int, row, col int) bool {
	for r, c := range cols {
		if c == col || r-c == row-col || r+c == row+col {
			return false
		}
	}
	return true
}

func drawChessBoard(cols []int) []string {
	res := make([]string, 0, len(cols))
	n := len(cols)
	for _, col := range cols {
		r := make([]byte, n)
		for i := 0; i < n; i++ {
			if i == col {
				r[i] = 'Q'
			} else {
				r[i] = '.'
			}
		}
		res = append(res, string(r))
	}
	return res
}

func solveNQueens_1(n int) [][]string {
	if n == 0 {
		return nil
	}

	var res [][]string
	chessboard := make([][]byte, n)
	for i := range chessboard {
		chessboard[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = '.'
		}
	}
	var dfs func(int, [][]byte)
	dfs = func(row int, chessboard [][]byte) {
		if row == n {
			board := make([]string, n)
			for i := range board {
				board[i] = string(chessboard[i])
			}
			res = append(res, board)
			return
		}

		for col := 0; col < n; col++ {
			if !isValid_1(row, col, chessboard) {
				continue
			}

			chessboard[row][col] = 'Q'
			dfs(row+1, chessboard)
			chessboard[row][col] = '.'
		}
	}

	dfs(0, chessboard)
	return res
}

func isValid_1(row, col int, chessboard [][]byte) bool {
	n := len(chessboard)
	for i := 0; i < row; i++ {
		if chessboard[i][col] == 'Q' {
			return false
		}
	}

	// 45 degree
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 135 degree
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
