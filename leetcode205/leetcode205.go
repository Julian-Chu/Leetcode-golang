package leetcode205

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
