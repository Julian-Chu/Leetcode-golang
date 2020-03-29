package leetcode205

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	set := make(map[byte]bool)

	sbytes := []byte(s)
	tbytes := []byte(t)

	for i := range sbytes {
		ch, ok := m[sbytes[i]]
		if !ok {
			if _, ok := set[tbytes[i]]; ok {
				return false
			}
			m[sbytes[i]] = tbytes[i]
			set[tbytes[i]] = true
			continue
		}

		if ch != tbytes[i] {
			return false
		}
	}

	return true
}
