package leetcode79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	var dfs func(int, int, int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || c < 0 || r > len(board)-1 || c > len(board[0])-1 || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = 0
		if dfs(r-1, c, index+1) || dfs(r+1, c, index+1) || dfs(r, c-1, index+1) || dfs(r, c+1, index+1) {
			return true
		}
		board[r][c] = temp
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
