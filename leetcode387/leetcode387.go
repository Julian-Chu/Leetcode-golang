package leetcode387

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	m := make(map[byte]int)
	b := []byte(s)

	for i := range b {
		m[b[i]]++
	}

	for i := range b {
		if m[b[i]] == 1 {
			return i
		}
	}
	return -1
}
