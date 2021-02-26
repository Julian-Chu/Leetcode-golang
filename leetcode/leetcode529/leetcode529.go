package leetcode529

func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	row, col := click[0], click[1]

	switch e := board[row][col]; {
	case e == 'M':
		board[row][col] = 'X'
		return board
	case e != 'E':
		return board
	}

	queue := [][]int{click}

	rowStart := 0
	rowEnd := len(board) - 1
	colStart := 0
	colEnd := len(board[0]) - 1
	for len(queue) > 0 {
		var nextq [][]int
		for _, e := range queue {

			row := e[0]
			col := e[1]
			if board[row][col] != 'E' {
				continue
			}
			matrix := [][]int{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1}, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			}

			mines := 0
			for _, m := range matrix {
				nextRow := row + m[0]
				nextCol := col + m[1]
				if (nextRow < rowStart) || (nextRow > rowEnd) || (nextCol < colStart) || (nextCol > colEnd) {
					continue
				}
				if board[nextRow][nextCol] == 'M' {
					mines++
				}
			}

			board[row][col] = 'B'
			if mines != 0 {
				board[row][col] = '0' + byte(mines)
			}
			//offsets := [][]int{
			//	{-1, 0}, {1, 0},
			//	{0, -1}, {0, 1},
			//}
			if board[row][col] == 'B' {
				for _, offset := range matrix {
					nextRow := row + offset[0]
					nextCol := col + offset[1]
					if (nextRow < rowStart) || (nextRow > rowEnd) || (nextCol < colStart) || (nextCol > colEnd) {
						continue
					}
					if board[nextRow][nextCol] == 'E' {
						nextq = append(nextq, []int{nextRow, nextCol})
					}
				}
			}
		}
		queue = nextq
	}

	return board
}
