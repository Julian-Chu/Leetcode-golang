package leetcode130

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	rows := len(board)
	cols := len(board[0])

	var handleSurround func(int, int)
	handleSurround = func(i, j int) {
		board[i][j] = 'm'
		if i > 0 && board[i-1][j] == 'O' {
			handleSurround(i-1, j)
		}

		if i < rows-1 && board[i+1][j] == 'O' {
			handleSurround(i+1, j)
		}

		if j > 0 && board[i][j-1] == 'O' {
			handleSurround(i, j-1)
		}

		if j < cols-1 && board[i][j+1] == 'O' {
			handleSurround(i, j+1)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
				if board[i][j] == 'O' {
					handleSurround(i, j)
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'm':
				board[i][j] = 'O'
			}
		}
	}
}
