package LeetCode_127_WordLadder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}

	delete(dict, beginWord)
	if dict[endWord] == false {
		return 0
	}

	q := map[string]bool{beginWord: true}
	length := 0
	notFound := true
loop:
	for len(q) > 0 {
		length++
		next := make(map[string]bool)

		for word := range q {
			w := []byte(word)
			for i := 0; i < len(w); i++ {
				ch := w[i]

				for j := 0; j < 26; j++ {
					newCh := 'a' + byte(j)
					if ch == newCh {
						continue
					}
					w[i] = newCh

					if dict[string(w)] == true {
						if string(w) == endWord {
							notFound = false
							break loop
						}
						next[string(w)] = true
						delete(dict, string(w))
					}
				}
				w[i] = ch
			}
		}
		q = next
	}

	if notFound {
		return 0
	}
	return length + 1

}

func ladderLength_2(beginWord string, endWord string, wordList []string) int {
	wordDict := map[string]struct{}{}

	for _, word := range wordList {
		wordDict[word] = struct{}{}
	}

	if _, ok := wordDict[endWord]; !ok {
		return 0
	}

	queue := []string{beginWord}
	visited := map[string]int{beginWord: 1}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			word := queue[i]
			path := visited[word]
			for j := 0; j < len(word); j++ {
				bs := []byte(word)
				for k := byte('a'); k <= byte('z'); k++ {
					bs[j] = k
					w := string(bs)
					if w == endWord {
						return path + 1
					}
					_, exist := wordDict[w]
					_, isVisited := visited[w]

					if exist && !isVisited {
						visited[w] = path + 1
						queue = append(queue, w)
					}
				}
			}
		}
		queue = queue[size:]
	}

	return 0
}
