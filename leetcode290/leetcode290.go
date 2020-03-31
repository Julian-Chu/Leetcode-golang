package leetcode290

import "strings"

func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0 {
		return false
	}

	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	m := make(map[byte]string)
	wordSet := make(map[string]bool)
	for i := range pattern {
		v, ok := m[pattern[i]]
		if !ok {
			if _, exist := wordSet[words[i]]; exist {
				return false
			}
			m[pattern[i]] = words[i]
			wordSet[words[i]] = true
			continue
		}

		if v != words[i] {
			return false
		}
	}
	return true

}
