package leetcode79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	val := false
	var dfs func([][]byte, int, int, int)
	dfs = func(boards [][]byte, r, c, index int) {
		if index == len(word) {
			val = true
			return
		}

		if r < 0 || c < 0 || r > len(board)-1 || c > len(board[0])-1 || board[r][c] != word[index] {
			return
		}

		temp := board[r][c]
		board[r][c] = 0
		dfs(board, r-1, c, index+1)
		dfs(board, r+1, c, index+1)
		dfs(board, r, c-1, index+1)
		dfs(board, r, c+1, index+1)
		board[r][c] = temp
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				dfs(board, i, j, 0)
			}
		}
	}
	return val
}
