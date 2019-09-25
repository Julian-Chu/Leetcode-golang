package leetcode79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	pointMap := make(map[Point]bool)
	var dfs func(map[Point]bool, []byte, Point)
	val := false
	dfs = func(mapper map[Point]bool, rest []byte, curPoint Point) {
		if len(rest) == 0 {
			val = true
			return
		}

		curP := curPoint

		w := rest[0]
		preRow := curP.RowIdx - 1
		if preRow >= 0 && board[preRow][curP.ColIdx] == w {
			nextPoint := Point{
				RowIdx: preRow,
				ColIdx: curP.ColIdx,
			}
			if _, ok := mapper[nextPoint]; !ok {
				mapper[nextPoint] = true
				dfs(mapper, rest[1:], nextPoint)
				delete(mapper, nextPoint)
			}
		}
		nextRow := curP.RowIdx + 1
		if nextRow < len(board) && board[nextRow][curP.ColIdx] == w {
			nextPoint := Point{
				RowIdx: nextRow,
				ColIdx: curP.ColIdx,
			}
			if _, ok := mapper[nextPoint]; !ok {
				mapper[nextPoint] = true
				dfs(mapper, rest[1:], nextPoint)
				delete(mapper, nextPoint)
			}
		}
		preCol := curP.ColIdx - 1
		if preCol >= 0 && board[curP.RowIdx][preCol] == w {
			nextPoint := Point{
				RowIdx: curP.RowIdx,
				ColIdx: preCol,
			}
			if _, ok := mapper[nextPoint]; !ok {
				mapper[nextPoint] = true
				dfs(mapper, rest[1:], nextPoint)
				delete(mapper, nextPoint)
			}
		}

		nextCol := curP.ColIdx + 1
		if nextCol < len(board[0]) && board[curP.RowIdx][nextCol] == w {
			nextPoint := Point{
				RowIdx: curP.RowIdx,
				ColIdx: nextCol,
			}
			if _, ok := mapper[nextPoint]; !ok {
				mapper[nextPoint] = true
				dfs(mapper, rest[1:], nextPoint)
				delete(mapper, nextPoint)
			}
		}
	}
	first := word[0]
	for i := range board {
		for j := range board[i] {
			if board[i][j] == first {
				curPoint := Point{
					RowIdx: i,
					ColIdx: j,
				}
				pointMap[curPoint] = true
				dfs(pointMap, []byte(word[1:]), curPoint)
				delete(pointMap, curPoint)
			}
		}
	}
	return val
}

type Point struct {
	RowIdx int
	ColIdx int
}
