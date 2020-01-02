package leetcode127

func ladderLength(beginWord string, endWord string, wordList []string) int {
}

func bfs(word, endWord string, cnt int, wordList []string) (newCnt int, newWordList []string) {

	for i := range wordList {
		if isOnlyOneDifferentAlphabet(word, wordList[i]) {

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
