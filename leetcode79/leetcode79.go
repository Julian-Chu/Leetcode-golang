package leetcode79

func exist(board [][]byte, word string) bool {
	boardmap := make(map[byte]int)
	wordmap := make(map[byte]int)

	for _, row := range board {
		for _, v := range row {
			boardmap[v]++
		}
	}

	for i := range word {
		w := word[i]
		if boardmap[w] == 0 || boardmap[w] == wordmap[w] {
			return false
		}
		wordmap[w]++
	}
	return true
}
