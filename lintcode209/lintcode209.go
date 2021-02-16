package lintcode209

/**
 * @param str: str: the given string
 * @return: char: the first unique character in a given string
 */
func firstUniqChar(str string) byte {
	m := make(map[byte]int)

	for _, ch := range str {
		c := byte(ch)
		if _, ok := m[c]; !ok {
			m[c] = 1
			continue
		}
		m[c]++
	}

	for ch, v := range m {
		if v == 1 {
			return ch
		}
	}
	return 0
}
