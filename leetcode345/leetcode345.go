package leetcode345

func reverseVowels(s string) string {
	bytes := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		if isVowel(bytes[l]) && isVowel(bytes[r]) {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
			continue
		}
		if !isVowel(bytes[l]) {
			l++
		}

		if !isVowel(bytes[r]) {
			r--
		}
	}

	return string(bytes)
}

func isVowel(ch byte) bool {
	if ch >= 'a' {
		ch -= 32
	}
	if ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}

	return false
}
