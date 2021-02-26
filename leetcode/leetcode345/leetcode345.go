package leetcode345

func reverseVowels(s string) string {
	b := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		if !isVowel(b[l]) {
			l++
			continue
		}

		if !isVowel(b[r]) {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	return string(b)
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'o', 'i', 'e', 'u', 'A', 'O', 'I', 'E', 'U':
		return true
	}

	return false
}
