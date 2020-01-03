package leetcode127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var q [][]string
	path := []string{beginWord}
	bfs(beginWord, endWord, path, wordList, &q)

	if len(q) == 0 {
		return 0
	}
	cnt := len(q[0])
	for i := range q {
		if len(q[i]) < cnt {
			cnt = len(q[i])
		}
	}
	return cnt
}

func bfs(word, endWord string, path, wordList []string, queue *[][]string) {
	if word == endWord {
		*queue = append(*queue, path)
		return
	}

	path = path[:len(path):len(path)]
	for i := range wordList {
		if isOnlyOneDifferentAlphabet(word, wordList[i]) {
			newPath := append(path, wordList[i])
			wordList[i], wordList[0] = wordList[0], wordList[i]
			newWordList := make([]string, len(wordList)-1)
			copy(newWordList, wordList[1:])
			bfs(wordList[0], endWord, newPath, newWordList, queue)
		}
	}
}

func isOnlyOneDifferentAlphabet(first, second string) bool {
	diff := 0
	size := len(first)
	for i := 0; i < size; i++ {
		if first[i] != second[i] {
			diff++
		}
	}
	return diff == 1
}
