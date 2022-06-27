package LeetCode_205_IsomorphicStrings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make([][]int, 265)

	for i := 0; i < len(m); i++ {
		m[i] = []int{-1, -1}
	}

	for i, sb := range s {
		tb := t[i]

		if m[sb][0] != m[tb][1] {
			return false
		}

		m[sb][0] = i
		m[tb][1] = i
	}

	return true
}

func isIsomorphic_2(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		if _, ok := s2t[s[i]]; !ok {
			s2t[s[i]] = t[i]
		}

		if _, ok := t2s[t[i]]; !ok {
			t2s[t[i]] = s[i]
		}

		if s2t[s[i]] != t[i] || t2s[t[i]] != s[i] {
			return false
		}
	}

	return true
}
