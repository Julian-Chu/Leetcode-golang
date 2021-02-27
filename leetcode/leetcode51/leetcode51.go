package leetcode51

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
