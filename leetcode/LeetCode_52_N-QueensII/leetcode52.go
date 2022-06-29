package LeetCode_52_N_QueensII

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	cols := make([]int, 0, n)
	return dfs(n, &cols)
}

func dfs(n int, cols *[]int) int {
	row := len(*cols)
	if row == n {
		return 1
	}

	sum := 0
	for col := 0; col < n; col++ {
		if !isValid(cols, row, col) {
			continue
		}
		*cols = append(*cols, col)
		sum += dfs(n, cols)
		*cols = (*cols)[:len(*cols)-1]
	}
	return sum
}

func isValid(cols *[]int, row, col int) bool {
	for r, c := range *cols {
		if c == col {
			return false
		}

		if r-c == row-col || r+c == row+col {
			return false
		}
	}
	return true
}

func totalNQueens_1(n int) int {
	res := 0

	chessboard := make([][]bool, n)
	for i := range chessboard {
		chessboard[i] = make([]bool, n)
	}

	var dfs func([][]bool, int)
	dfs = func(board [][]bool, row int) {
		if row == len(board) {
			res++
			return
		}

		for i := range board[row] {
			if isValid_1(board, row, i) {
				board[row][i] = true
				dfs(board, row+1)
				board[row][i] = false
			}
		}
	}
	dfs(chessboard, 0)
	return res
}

func isValid_1(board [][]bool, row, col int) bool {
	for r := row; r >= 0; r-- {
		if r-row+col >= 0 && board[r][r-row+col] {
			return false
		}
		if board[r][col] {
			return false
		}
		if row+col-r < len(board) && board[r][row+col-r] {
			return false
		}
	}

	return true
}
