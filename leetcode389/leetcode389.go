package leetcode389

func findTheDifference(s string, t string) byte {
	sb := []byte(s)
	tb := []byte(t)

	m := make(map[byte]int)

	for _, b := range sb {
		m[b]++
	}

	for _, b := range tb {
		if _, ok := m[b]; !ok {
			return b
		}
		m[b]--

	}

	for key, cnt := range m {
		if cnt != 0 {
			return key
		}

	}

	return 0
}
