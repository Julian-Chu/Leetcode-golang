package leetcode720

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	res := ""
	set := make(map[string]bool, len(words))

	for _, word := range words {
		if len(word) > 1 && !set[word[:len(word)-1]] {
			continue
		}

		set[word] = true

		if len(word) > len(res) {
			res = word
		}
	}
	return res
}
