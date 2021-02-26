package leetcode392

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	cnt := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			cnt++
		}
		j++
	}

	return cnt == len(s)
}
