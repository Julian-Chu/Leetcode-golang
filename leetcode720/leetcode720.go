package leetcode720

import "sort"

func longestWord(words []string) string {
	m := make(map[string]bool)

	for _, word := range words {
		m[word] = true
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) > len(words[j])
	})

	for _, word := range words {
		for i := range word {
			if _, ok := m[word[:i+1]]; !ok {
				break
			}
			if i == len(word)-1 {
				return word
			}
		}
	}
	return ""
}
