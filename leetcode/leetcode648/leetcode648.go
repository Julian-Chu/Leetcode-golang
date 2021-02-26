package leetcode648

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	hasRoot := make(map[string]bool, len(dictionary))
	hasLen := make([]bool, 101)

	for i := range dictionary {
		hasRoot[dictionary[i]] = true
		hasLen[len(dictionary[i])] = true
	}

	ls := make([]int, 0, 101)
	for i, ok := range hasLen {
		if ok {
			ls = append(ls, i)
		}
	}
	sort.Ints(ls)

	words := strings.Split(sentence, " ")
	for i, w := range words {
		for _, l := range ls {
			if l < len(w) && hasRoot[w[:l]] {
				words[i] = w[:l]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
