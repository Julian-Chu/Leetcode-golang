package leetcode127

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
