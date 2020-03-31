package leetcode290

import "strings"

func wordPattern(pattern string, str string) bool {
	words := strings.Fields(str)
	if len(pattern) != len(words) {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		b := pattern[i]
		s := words[i]

		bs, ok1 := m1[b]
		sb, ok2 := m2[s]

		if !ok1 && !ok2 {
			m1[b] = s
			m2[s] = b
		}

		if ok1 && ok2 && (b != sb || s != bs) {
			return false
		}

		if (ok1 && !ok2) || (!ok1 && ok2) {
			return false
		}
	}
	return true
}
