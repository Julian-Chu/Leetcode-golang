package leetcode289

func gameOfLife(board [][]int) {
	for i := range board {
		for j := range board[i] {
			board[i][j] = getNextState(board, i, j)
		}
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] %= 2
		}
	}
}

func getNextState(board [][]int, i int, j int) int {
	rows := len(board)
	cols := len(board[i])

	rowStart := i - 1
	if i-1 < 0 {
		rowStart = i
	}
	rowEnd := i + 1
	if i+1 > rows-1 {
		rowEnd = i
	}
	colStart := j - 1
	if j-1 < 0 {
		colStart = j
	}

	colEnd := j + 1
	if j+1 > cols-1 {
		colEnd = j
	}

	cnt := 0
	for m := rowStart; m <= rowEnd; m++ {
		for n := colStart; n <= colEnd; n++ {
			if m == i && n == j {
				continue
			}
			if board[m][n] == 1 || board[m][n] == 2 {
				cnt++
			}
		}
	}
	if board[i][j] == 1 && (cnt < 2 || cnt > 3) {
		return 2
	} else if board[i][j] == 0 && cnt == 3 {
		return 3
	}
	return board[i][j]
}
