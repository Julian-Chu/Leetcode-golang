package LeetCode_459_RepeatedSubstringPattern

func repeatedSubstringPattern_KMP(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		sub := s[:i+1]
		if isPattern(sub, s) {
			return true
		}
	}
	return false
}

func isPattern(sub, s string) bool {
	l := len(sub)
	for i := 0; i < len(s); {
		// when rest characters of string is less than
		if i+l > len(s) {
			return false
		}
		if s[i:i+l] != sub {
			return false
		}

		i += l
	}
	return true
}
