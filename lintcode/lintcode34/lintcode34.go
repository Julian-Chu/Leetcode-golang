package lintcode34

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
