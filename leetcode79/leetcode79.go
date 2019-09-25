package leetcode79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	var dfs func([]Point, []byte)
	val := false
	dfs = func(path []Point, rest []byte) {
		path = path[:len(path):len(path)]
		if len(rest) == 0 {
			val = true
			return
		}

		curP := path[len(path)-1]

		w := rest[0]
		preRow := curP.RowIdx - 1
		if preRow >= 0 && board[preRow][curP.ColIdx] == w {
			point := Point{
				RowIdx: preRow,
				ColIdx: curP.ColIdx,
			}
			if isInPath(path, point) == false {
				dfs(append(path, point), rest[1:])
			}
		}
		nextRow := curP.RowIdx + 1
		if nextRow < len(board) && board[nextRow][curP.ColIdx] == w {
			point := Point{
				RowIdx: nextRow,
				ColIdx: curP.ColIdx,
			}
			if isInPath(path, point) == false {
				dfs(append(path, point), rest[1:])
			}
		}
		preCol := curP.ColIdx - 1
		if preCol >= 0 && board[curP.RowIdx][preCol] == w {
			point := Point{
				RowIdx: curP.RowIdx,
				ColIdx: preCol,
			}
			if isInPath(path, point) == false {
				dfs(append(path, point), rest[1:])
			}
		}

		nextCol := curP.ColIdx + 1
		if nextCol < len(board[0]) && board[curP.RowIdx][nextCol] == w {
			point := Point{
				RowIdx: curP.RowIdx,
				ColIdx: nextCol,
			}
			if isInPath(path, point) == false {
				dfs(append(path, point), rest[1:])
			}
		}
	}
	first := word[0]
	for i := range board {
		for j := range board[i] {
			if board[i][j] == first {
				dfs([]Point{{
					RowIdx: i,
					ColIdx: j,
				}}, []byte(word[1:]))
			}
		}
	}
	return val
}

type Point struct {
	RowIdx int
	ColIdx int
}

func isInPath(path []Point, point Point) bool {
	for _, p := range path {
		if point.ColIdx == p.ColIdx && point.RowIdx == p.RowIdx {
			return true
		}
	}
	return false
}
