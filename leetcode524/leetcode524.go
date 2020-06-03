package leetcode524

import "sort"

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) == len(d[j]) {
			return d[i] < d[j]
		}

		return len(d[i]) > len(d[j])
	})

	for _, word := range d {
		if isSub(word, s) {
			return word
		}
	}

	return ""
}

func isSub(word string, s string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(s) {
		if word[i] == s[j] {
			i++
		}
		j++
	}
	return i == len(word)
}
